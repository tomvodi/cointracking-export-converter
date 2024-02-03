<script lang="ts" setup>
import AddExportFile from "./components/AddExportFile.vue";
import ExportFilesList from "./components/ExportFilesList.vue";</script>

<template>
  <v-app>
    <v-app-bar :elevation="2">
      <v-app-bar-title>Converter</v-app-bar-title>
    </v-app-bar>

    <v-main>
      <v-container class="mb-6 my-3">
        <AddExportFile/>
        <ExportFilesList/>
      </v-container>
    </v-main>
  </v-app>
</template>

<style>
</style>
