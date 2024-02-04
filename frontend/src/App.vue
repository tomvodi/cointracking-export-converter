<script lang="ts" setup>
import AddExportFile from "./components/AddExportFile.vue";
import ExportFilesList from "./components/ExportFilesList.vue";</script>

<template>
  <v-app>
    <v-app-bar :elevation="2" color="blue-grey-lighten-5">
      <v-app-bar-title>CoinTracking CSV Export File Converter</v-app-bar-title>
    </v-app-bar>
    <v-main>
      <v-container class="my-4 pa-3">
        <AddExportFile class="mb-4"/>
        <ExportFilesList/>
      </v-container>
    </v-main>
  </v-app>
</template>

<style>
</style>
