<script setup lang="ts">
import { ref } from 'vue'
import {OpenExportFile} from "../../wailsjs/go/cointracking/ct";

const selectedFile = ref<string>('')

const selectFile = async () => {
OpenExportFile().then((value:string)=> {
  console.log("file was selected: " + value)
  selectedFile.value = value
}).catch((reason:any) => {
  console.log("error selecting file: " + reason)
})
}

const startConvertFile = async () => {
  if (!selectedFile.value?.length) {
    return
  }
}
</script>

<template>
  <v-row>
    <v-text-field clearable label="Cointracking export file (csv)"
                  variant="outlined" readonly
                  v-model="selectedFile"></v-text-field>
    <v-btn @click="selectFile">Select File</v-btn>
  </v-row>
  <v-row>
    <v-btn @click="startConvertFile">Convert</v-btn>
  </v-row>
</template>

<style scoped>

</style>