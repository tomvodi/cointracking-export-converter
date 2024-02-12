<script setup lang="ts">

import {ref} from "vue";
import InfoButton from "./InfoButton.vue";
import {useSettingsStore} from "../stores/SettingsStore";

const props = defineProps({
  selectedTimezone: String
})

const rawSelected = ref(props.selectedTimezone)
const store = useSettingsStore()

const $emit = defineEmits<{
  timezoneChanged: [zone: string]
}>()

const onTimezoneSelected = (newTimezone: string) => {
  console.log("timezone selected " + newTimezone)
  $emit('timezoneChanged', newTimezone)
}
</script>

<template>
  <v-autocomplete
      :items="store.allTimezones"
      v-model="rawSelected"
      label="Transaction Timezone"
      @update:model-value="onTimezoneSelected"
  >
    <template #append>
      <InfoButton>
        <div class="text-subtitle-2 font-weight-bold pb-1">Transaction Timezone</div>
        <div class="text-body-2">CoinTracking exports transactions with timestamps according to the timezone set in
          <span class="font-weight-bold">Account > Account Settings > Your time zone</span>.<br/>
          This means that when a
          transaction happens at a specific point in time, the exported date and time will look different for someone
          who set the timezone to Europe/Lisbon and someone who set it to Australia/Perth for example. <br/>
          This deviation of many hours will make it hard to match outgoing and incoming transactions for blockpit, so
          you have to set here exactly the same timezone that you selected in CoinTracking to have valid data in the
          end.
        </div>
      </InfoButton>
    </template>
  </v-autocomplete>
</template>

<style scoped>

</style>