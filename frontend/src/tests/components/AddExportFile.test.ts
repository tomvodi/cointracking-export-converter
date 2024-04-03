import {describe, expect, test, vi} from "vitest";
import {createVuetify} from "vuetify";
import * as components from 'vuetify/components'
import * as directives from 'vuetify/directives'
import {mount, RouterLinkStub, VueWrapper} from "@vue/test-utils";
import AddExportFile from "@/components/AddExportFile.vue";
import {createTestingPinia} from "@pinia/testing";
import {useSettingsStore} from "@/stores/settingsStore";
import {router} from "@/router";

const vuetify = createVuetify({
    components,
    directives,
})

const pinia = createTestingPinia({
    initialState: {
        settings: {
            timezone: "",
        }
    }
})

global.ResizeObserver = require('resize-observer-polyfill')

function defaultWrapper(): VueWrapper {
    const wrapper = mount(AddExportFile, {
        global: {
            plugins: [
                router,
                pinia,
                vuetify,
            ]
        },
        stubs: {
            RouterLink: RouterLinkStub,
        },
    });

    return wrapper;
}

describe("initial state", () => {
    test("should show notification when no timezone is set initially", async () => {
        const store = vi.mocked(useSettingsStore(pinia));

        // @ts-ignore
        store.timezoneEmpty = true;

        const wrapper = defaultWrapper();
        expect(wrapper.find('.v-alert').exists()).toBeTruthy();
        const selectFileBtn = wrapper.find('.v-btn');
        expect(selectFileBtn.attributes().disabled).toBeDefined();
    });
})

describe("timezone set", () => {
    test("should show everything as expected", async () => {
        const store = useSettingsStore(pinia);
        // @ts-ignore
        store.timezoneEmpty = false;
        const wrapper = defaultWrapper();
        expect(wrapper.find('.v-alert').exists()).toBeFalsy();
        const selectFileBtn = wrapper.find('.v-btn');
        expect(selectFileBtn.attributes().disabled).toBeDefined();
    })
})

describe("file will be opened", () => {
    test("should open file", async () => {
        const wrapper = defaultWrapper();
        const selectFileBtn = wrapper.find('.v-btn');
        await selectFileBtn.trigger('click');
        expect(wrapper.emitted('selectFile')).toBeTruthy();
    })
})