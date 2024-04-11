import {beforeEach, describe, expect, test} from "vitest";
import {createPinia, setActivePinia} from "pinia";
import {useSnackbarStore} from "@/stores/snackbar_store";

describe("snackbar store", () => {
    beforeEach(() => {
        setActivePinia(createPinia())
    })

    test("showError", () => {
        const store = useSnackbarStore()
        store.showError("test error")
        expect(store.visible).toBe(true)
        expect(store.text).toBe("test error")
        expect(store.color).toBe("red-lighten-1")
    })

    test("hideSnackbar", () => {
        const store = useSnackbarStore()
        store.showError("test error")
        store.hideSnackbar()
        expect(store.visible).toBe(false)
    })
})