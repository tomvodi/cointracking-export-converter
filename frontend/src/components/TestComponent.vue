<script setup lang="ts">
import {inject, onMounted} from "vue";
import {EventsOn} from "@wails/runtime";
import {common} from "@wails/go/models";
import {WailsApi} from "@/wails/wails_api";
import {wailsClientInjKey} from "@/injection_keys";

const wailsClient: WailsApi = inject<WailsApi>(wailsClientInjKey) as WailsApi

onMounted(() => {
  EventsOn("ExportFilesChanged", setExportFiles)
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