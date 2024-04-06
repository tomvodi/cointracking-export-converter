import {defineStore} from "pinia";
import {common} from "@wails/go/models";
import {computed, ref} from "vue";

export const useApplicationStore = defineStore('application', () => {
    const exportFiles = ref<Array<common.ExportFileInfo>>([])

    const hasExportFiles = computed(() => {
        return exportFiles.value.length > 0
    })

    return {exportFiles, hasExportFiles}
});
