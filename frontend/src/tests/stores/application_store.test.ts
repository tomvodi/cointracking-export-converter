import {beforeEach, describe, expect, test} from "vitest";
import {createPinia, setActivePinia} from "pinia";
import {useApplicationStore} from "@/stores/application_store";
import {common} from "@wails/go/models";

describe("application store", () => {
    beforeEach(() => {
        setActivePinia(createPinia())
    })

    test("hasExportFiles", () => {
        const store = useApplicationStore()
        expect(store.hasExportFiles).toBe(false)

        store.exportFiles.push({fileName: "test"} as common.ExportFileInfo)
        expect(store.hasExportFiles).toBe(true)
    })
})