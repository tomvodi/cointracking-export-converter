import {createVuetify} from "vuetify";
import * as components from 'vuetify/components'
import * as directives from 'vuetify/directives'
import {describe, expect, test, vi} from "vitest";
import {createTestingPinia} from "@pinia/testing";
import {mount} from "@vue/test-utils";
import ExportView from "@/views/ExportView.vue";
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


function defaultWrapper() {
    const mockedEventsOn = vi.fn((event: string, callback: Function) => {
        console.log("mockedEventsOn called");
    })
    const mockedEventsOnMultiple = vi.fn((event: string, callback: Function, maxCallback: number) => {
        console.log("mockedEventsOnMultiple called");
    })
    const wrapper = mount(ExportView, {
        global: {
            plugins: [
                router,
                pinia,
                vuetify,
            ],
            mocks: {
                EventsOn: mockedEventsOn,
                EventsOnMultiple: mockedEventsOnMultiple,
            },
        },
    });

    return wrapper;
}

describe("file is selected", () => {
    test("should call OpenExportFile on wails interface", async () => {
        const mockedEventsOn = vi.fn((event: string, callback: Function) => {
            console.log("mockedEventsOn called");
        })
        global.EventsOn = mockedEventsOn;
        const mockedEventsOnMultiple = vi.fn((event: string, callback: Function, maxCallback: number) => {
            console.log("mockedEventsOnMultiple called");
        })
        global.EventsOnMultiple = mockedEventsOnMultiple;
        const mockOpenExport = vi.fn((tz: string) => {
            return Promise.resolve("file opened");
        })
        global.OpenExportFile = mockOpenExport;
        const wrapper = defaultWrapper();
        const exportFile = wrapper.find('AddExportFile');
        await exportFile.trigger('selectFile')
        await wrapper.vm.$nextTick();

        expect(mockOpenExport).toHaveBeenCalled();
    });
})
