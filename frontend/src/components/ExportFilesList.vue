<script setup lang="ts">
import {useApplicationStore} from "@/stores/application_store";

const appStore = useApplicationStore();
import {common} from "@wails/go/models";
import {WailsRuntimeApi} from "@/wails/wails_runtime_api";
import {wailsClientInjKey, wailsRuntimeInjKey} from "@/injection_keys";
import {WailsApi} from "@/wails/wails_api";

const wailsRuntime: WailsRuntimeApi = inject<WailsRuntimeApi>(wailsRuntimeInjKey) as WailsRuntimeApi
const wailsClient: WailsApi = inject<WailsApi>(wailsClientInjKey) as WailsApi

const $emit = defineEmits<{
  exportFilesChanged: [files: Array<common.ExportFileInfo>],
}>()

const exportedFiles = ref<Array<common.ExportFileInfo>>([])

onMounted(() => {
  wailsRuntime.EventsOn("ExportFilesChanged", setExportFiles)

  wailsClient.GetExportFiles().then(setExportFiles).catch((reason: any) => {
    console.log("failed getting export files initially: " + reason)
  })
})

const setExportFiles = async (files: Array<common.ExportFileInfo>) => {
  exportedFiles.value = files
  $emit('exportFilesChanged', exportedFiles.value)
}

</script>

<template>
  <div class="px-5" v-if="appStore.hasExportFiles">
    <p class="text-h6">CoinTracking export files</p>
    <v-list lines="two">
      <v-list-item
          v-for="exportFile in appStore.exportFiles"
          :title="exportFile.fileName"
          :subtitle="`${exportFile.txCount} transactions on ${exportFile.exchanges.length} exchanges. ${exportFile.skippedTxs} skipped transactions.`"
      >
        <template v-slot:prepend>
          <v-icon icon="mdi-file-delimited-outline"></v-icon>
        </template>
      </v-list-item>
    </v-list>
  </div>
</template>

<style scoped>

</style>