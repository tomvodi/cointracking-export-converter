import {createPinia, setActivePinia} from "pinia"
import {useSettingsStore} from "@/stores/settings_store"
import {beforeEach, describe, expect, test} from "vitest"

describe("settings store", () => {
    beforeEach(() => {
        setActivePinia(createPinia())
    })

    test("timezoneEmpty", () => {
        const store = useSettingsStore()
        expect(store.timezoneEmpty).toBeTruthy()

        store.settingsLoaded = true
        expect(store.timezoneEmpty).toBe(true)

        store.timezone = "Europe/Lisabon"
        expect(store.timezoneEmpty).toBe(false)
    })
})