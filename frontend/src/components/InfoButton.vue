<script setup lang="ts">
import {ref} from "vue";

const show = ref(false)

</script>

<template>
  <v-row class="ma-1 pa-1">
    <v-tooltip v-model="show" location="left" max-width="500">
      <template v-slot:activator="{ props }">
        <v-icon v-bind="props" icon="mdi-alert-circle"></v-icon>
      </template>
      <v-sheet class="pa-2 text-left" color="info" :elevation="8" rounded>
        <slot></slot>
      </v-sheet>
    </v-tooltip>
  </v-row>

</template>

<style scoped>
:deep(.v-overlay__content) {
  background: rgba(255, 66, 66, 0) !important;
}
</style>