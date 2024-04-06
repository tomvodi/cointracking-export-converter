<script setup lang="ts">
import {useApplicationStore} from "@/stores/application_store";

const appStore = useApplicationStore();
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