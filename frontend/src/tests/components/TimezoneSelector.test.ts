import {describe, expect, test, vi} from "vitest";
import {createVuetify} from 'vuetify'
import * as components from 'vuetify/components'
import * as directives from 'vuetify/directives'
import TimezoneSelector from "@/components/TimezoneSelector.vue";
import {mount, VueWrapper} from "@vue/test-utils";
import {createTestingPinia} from "@pinia/testing";
import {useSettingsStore} from "@/stores/settingsStore";

const vuetify = createVuetify({
    components,
    directives,
})

global.ResizeObserver = require('resize-observer-polyfill')

function defaultWrapper(): VueWrapper {
    const wrapper = mount(TimezoneSelector, {
        global: {
            plugins: [
                createTestingPinia({
                    createSpy: vi.fn,
                    initialState: {
                        settings: {
                            allTimezones: [
                                {value: "america_new_york", title: "America/New_York"},
                                {value: "america_los_angeles", title: "America/Los_Angeles"},
                                {value: "europe_london", title: "Europe/London"},
                            ]
                        }
                    }
                }),
                vuetify,
            ]
        },
        props: {selectedTimezone: "america_los_angeles"},
    });

    return wrapper;
}

describe("initial selector state", () => {
    test("should have initial values", async () => {
        const wrapper = defaultWrapper();
        const store = useSettingsStore();
        const timezoneSelect = wrapper.getComponent({name: 'v-autocomplete'});
        expect(timezoneSelect.props('items')).toEqual(store.allTimezones);
    });
})

describe("selecting a timezone", () => {
    test("should emit changed timezone", async () => {
        const wrapper = defaultWrapper();
        const timezoneSelect = wrapper.getComponent({name: 'v-autocomplete'});
        await timezoneSelect.setValue("europe_london");
        await wrapper.vm.$nextTick();
        const tzChangeEvent = wrapper.emitted('timezoneChanged')
        expect(tzChangeEvent).toHaveLength(1);
        if (tzChangeEvent) {
            expect(tzChangeEvent[0][0]).toEqual("europe_london");
        }
    });
})