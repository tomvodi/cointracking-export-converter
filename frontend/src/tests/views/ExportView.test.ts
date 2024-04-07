import {createVuetify} from "vuetify";
import * as components from 'vuetify/components'
import * as directives from 'vuetify/directives'
import {describe, expect, test, vi} from "vitest";
import {createTestingPinia} from "@pinia/testing";
import ExportView from "@/views/ExportView.vue";
import {router} from "@/router";
import {flushPromises, mount} from "@vue/test-utils";
import {useSnackbarStore} from "@/stores/snackbar_store";
import {useSettingsStore} from "@/stores/settings_store";

const vuetify = createVuetify({
    components,
    directives,
})

const testTimezone = "Europe/Lisabon"
const openExportFileMock = vi.fn()
const exportToBlockpitMock = vi.fn()

const pinia = createTestingPinia({
    createSpy: vi.fn,
    initialState: {
        settings: {
            timezone: testTimezone,
        },
        snackbar: {
            visible: false,
            text: '',
            color: '',
            timeout: 5000,
        },
    }
})

global.ResizeObserver = require('resize-observer-polyfill')

function defaultWrapper() {
    const wrapper = mount(ExportView, {
        global: {
            plugins: [
                router,
                pinia,
                vuetify,
            ],
            stubs: {
                ExportFilesList: {
                    template: '<div />',
                }
            },
            provide: {
                wailsClient: {
                    OpenExportFile: openExportFileMock,
                    ExportToBlockpitXlsx: exportToBlockpitMock,
                },
            }
        },
    });

    return wrapper;
}

describe("selecting a file", () => {
    test("should call OpenExportFile on wails interface", async () => {
        openExportFileMock.mockReturnValue(Promise.resolve());

        const wrapper = defaultWrapper();
        const exportFile = wrapper.getComponent({name: 'AddExportFile'});
        await exportFile.vm.$emit('selectFile');

        expect(openExportFileMock).toHaveBeenCalledWith(testTimezone);
    });

    test("should show notification when OpenExportFile fails", async () => {
        const errMsg = 'failed to open file';
        openExportFileMock.mockReturnValue(Promise.reject(errMsg));

        const snackStore = vi.mocked(useSnackbarStore(pinia));

        const wrapper = defaultWrapper();
        const exportFile = wrapper.getComponent({name: 'AddExportFile'});
        await exportFile.vm.$emit('selectFile');
        await flushPromises();

        expect(snackStore.showError).toHaveBeenCalledWith(errMsg);
    });
})

describe("exporting to Blockpit", () => {
    test("should disable button if there are no export files", async () => {
        const store = useSettingsStore(pinia);
        // @ts-ignore
        store.hasExportFiles = false
        const wrapper = defaultWrapper();

        expect(wrapper.find('.v-btn').exists()).toBeTruthy();
        const saveButton = wrapper.find('.v-btn');
        expect(saveButton.attributes().disabled).toBeDefined();
    })

    test("should enable button if there are no export files", async () => {
        const store = useSettingsStore(pinia);
        // @ts-ignore
        store.hasExportFiles = true
        const wrapper = defaultWrapper();

        expect(wrapper.find('.v-btn').exists()).toBeTruthy();
        const saveButton = wrapper.find('.v-btn');
        expect(saveButton.attributes().disabled).toBeFalsy();
    })

    test("should export to blockpit", async () => {
        exportToBlockpitMock.mockReturnValue(Promise.resolve());

        const wrapper = defaultWrapper();
        const saveButton = wrapper.findComponent({ref: 'saveBpBtn'});
        expect(saveButton.exists()).toBeTruthy();
        await saveButton.vm.$emit('click');
        await flushPromises();

        expect(exportToBlockpitMock).toHaveBeenCalled();
    })

    test("should show error when export failed", async () => {
        const errMsg = 'failed to export file';
        exportToBlockpitMock.mockReturnValue(Promise.reject(errMsg));

        const snackStore = vi.mocked(useSnackbarStore(pinia));

        const wrapper = defaultWrapper();
        const saveButton = wrapper.findComponent({ref: 'saveBpBtn'});
        expect(saveButton.exists()).toBeTruthy();
        await saveButton.vm.$emit('click');
        await flushPromises();

        expect(snackStore.showError).toHaveBeenCalledWith(`failed saving blockpit file: ${errMsg}`);
    })
})
