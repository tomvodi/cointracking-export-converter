<script setup lang="ts">
import {inject, onMounted} from "vue";
import {common} from "@wails/go/models";
import {WailsApi} from "@/wails/wails_api";
import {wailsClientInjKey, wailsRuntimeInjKey} from "@/injection_keys";
import {WailsRuntimeApi} from "@/wails/wails_runtime_api";

const wailsClient: WailsApi = inject<WailsApi>(wailsClientInjKey) as WailsApi
const wailsRuntime: WailsRuntimeApi = inject<WailsRuntimeApi>(wailsRuntimeInjKey) as WailsRuntimeApi

onMounted(() => {
  wailsRuntime.EventsOn("ExportFilesChanged", setExportFiles)
})

const setExportFiles = async (files: Array<common.ExportFileInfo>) => {
  console.log("files", files)
}

const selectFile = async () => {
  wailsClient.OpenExportFile("Europe/Lisabon").catch((reason: any) => {
    console.log("reason", reason)
  })
}

</script>

<template>
  <v-btn @click="selectFile">Select</v-btn>
</template>

<style scoped>

</style>