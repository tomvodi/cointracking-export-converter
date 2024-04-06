<script setup lang="ts">
import {useSnackbarStore} from "@/stores/snackbar_store";

const store = useSnackbarStore()

</script>
<template>
  <v-snackbar
      v-model="store.visible"
      :timeout="store.timeout"
      :color="store.color"
  >
    {{ store.text }}
    <template v-slot:actions>
      <v-btn
          color="black"
          variant="text"
          @click="store.hideSnackbar()"
      >
        Close
      </v-btn>
    </template>
  </v-snackbar>
</template>

<style scoped>

</style>