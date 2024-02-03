<script setup lang="ts">

import {EventsOn} from "../../wailsjs/runtime";
import {onMounted, ref} from "vue";

interface ExportFile {
  filePath: string
  txCount: number
  exchanges: Array<string>
}

const exportedFiles = ref<Array<ExportFile>>([])


onMounted(() => {
  EventsOn("ExportFilesChanged", (files: Array<ExportFile>) => {
    console.log("ExportFilesChanged")
    exportedFiles.value = files
  })
})

</script>

<template>
  <v-list lines="two">
    <v-list-subheader inset>CoinTracking export files</v-list-subheader>
    <v-list-item
        v-for="exportFile in exportedFiles"
        :title="exportFile.filePath"
        :subtitle="`${exportFile.txCount} transactions on ${exportFile.exchanges.length} exchanges`"
    />
  </v-list>
</template>

<style scoped>

</style>