<script setup lang="ts">

import {EventsOn} from "../../wailsjs/runtime";
import {onMounted, ref} from "vue";

interface ExportFile {
  fileName: string
  txCount: number
  exchanges: Array<string>
}

const exportedFiles = ref<Array<ExportFile>>([])


onMounted(() => {
  EventsOn("ExportFilesChanged", (files: Array<ExportFile>) => {
    exportedFiles.value = files
  })
})

</script>

<template>
  <div class="px-5" v-if="exportedFiles.length > 0">
    <p class="text-h6">CoinTracking export files</p>
    <v-list lines="two">
      <v-list-item
          v-for="exportFile in exportedFiles"
          :title="exportFile.fileName"
          :subtitle="`${exportFile.txCount} transactions on ${exportFile.exchanges.length} exchanges`"
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