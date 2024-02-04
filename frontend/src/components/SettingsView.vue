<script setup lang="ts">

import {useRouter} from "vue-router";
import TitledPanel from "./TitledPanel.vue";

const router = useRouter()

const navigateBack = () => {
  router.back()
}

</script>

<template>
  <TitledPanel title="Settings">
    <template v-slot:append>
      <v-btn
          variant="plain"
          @click="navigateBack"
          icon="mdi-close-circle-outline"
      ></v-btn>
    </template>
  </TitledPanel>
</template>

<style scoped>

</style>