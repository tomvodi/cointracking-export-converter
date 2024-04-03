import {describe, expect, test, vi} from "vitest";
import {mount} from "@vue/test-utils";
import TestComponent from "./TestComponent.vue";
import {createVuetify} from "vuetify";
import * as components from "vuetify/components";
import * as directives from "vuetify/directives";

const vuetify = createVuetify({
    components,
    directives,
})

describe("TestComponent", () => {
    test("should call OpenExportFile on wails interface", async () => {
        const openExportMock = vi.fn()
        const eventsOnMock = vi.fn()

        const wrapper = mount(TestComponent, {
            global: {
                plugins: [
                    vuetify,
                ],
                provide: {
                    wailsClient: {
                        OpenExportFile: openExportMock,
                    },
                    wailsRuntime: {
                        EventsOn: eventsOnMock,
                    }
                }
            },
        })

        openExportMock.mockReturnValueOnce(Promise.reject('failed to open file'));
        const button = wrapper.find('.v-btn');
        expect(button.exists()).toBeTruthy();
        await button.trigger('click');
        expect(openExportMock).toHaveBeenCalledWith('Europe/Lisabon');
        expect(eventsOnMock).toHaveBeenCalledWith('ExportFilesChanged', expect.any(Function));
    });
})