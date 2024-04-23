import {createVuetify} from "vuetify";
import * as components from 'vuetify/components'
import * as directives from 'vuetify/directives'
import {describe, expect, test, vi} from "vitest";
import {createTestingPinia} from "@pinia/testing";
import {router} from "@/router";
import {mount} from "@vue/test-utils";
import {useSettingsStore} from "@/stores/settings_store";
import SettingsView from "@/views/SettingsView.vue";

const vuetify = createVuetify({
    components,
    directives,
})

const defaultTimezone = "Europe/Amsterdam"
const testTimezone = "Europe/Lisabon"
const wailsSetTimezone = vi.fn()
const pinia = createTestingPinia({
    createSpy: vi.fn,
    initialState: {
        settings: {
            timezone: defaultTimezone,
            allTimezones: [
                {value: 'default', title: 'Europe/Amsterdam'},
                {value: 'new', title: 'Europe/Lisabon'},
            ],
        },
    }
})

global.ResizeObserver = require('resize-observer-polyfill')

function defaultWrapper() {
    const wrapper = mount(SettingsView, {
        global: {
            plugins: [
                router,
                pinia,
                vuetify,
            ],
            provide: {
                wailsClient: {
                    SetTimezone: wailsSetTimezone,
                },
            }
        },
    });

    return wrapper;
}

describe('change timezone', () => {
    test('change timezone', async () => {
        const wrapper = defaultWrapper();

        const settingsStore = useSettingsStore();
        expect(settingsStore.timezone).toBe(defaultTimezone);

        const tzSelector = wrapper.findComponent({ref: 'tzSelector'})
        expect(tzSelector.exists()).toBeTruthy();
        tzSelector.vm.$emit('timezoneChanged', testTimezone);
        expect(settingsStore.timezone).toBe(testTimezone);
    })
})