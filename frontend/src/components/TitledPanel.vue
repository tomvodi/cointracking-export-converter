<script setup lang="ts">
const props = defineProps({
  title: String
})
</script>

<template>
  <v-sheet class="pa-5" border>
    <v-row>
      <p class="text-h5 mt-2 ml-3 pb-7">{{ props.title }}</p>
      <v-spacer></v-spacer>
      <slot name="append"></slot>
    </v-row>
    <slot></slot>

  </v-sheet>
</template>

<style scoped>

</style>