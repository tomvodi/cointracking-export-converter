<script setup lang="ts">

import {useRouter} from "vue-router";
import {onMounted, ref} from "vue";
import {
  BlockpitTxTypes,
  SetCointracking2BlockpitMapping,
  SetTimezone,
  Timezone,
  TxTypeMappings
} from "../../wailsjs/go/config/appConfig";
import {common} from "../../wailsjs/go/models";
import {useSettingsStore} from "../stores/SettingsStore";
import TitledPanel from "../components/TitledPanel.vue";
import TimezoneSelector from "../components/TimezoneSelector.vue";

const router = useRouter()
const store = useSettingsStore()

const txMappings = ref<Array<common.Ct2BpTxMapping>>()
let blockpitTxTypes = ref(Array<common.TxDisplayName>())

const navigateBack = () => {
  router.back()
}

onMounted(() => {
  TxTypeMappings().then((mappings) => {
    txMappings.value = mappings
  }).catch((reason) => {
    console.log(`failed getting tx type mappings between cointracking and blockpit: ${reason}`)
  })

  BlockpitTxTypes().then((displayNames) => {
    blockpitTxTypes.value = displayNames
  }).catch((reason) => {
    console.log(`failed getting blockpit tx types: ${reason}`)
  })

  Timezone().then((timezone: string) => {
    store.timezone = timezone
  })
})

const blockpitTxTypeChanged = (idx: number) => {
  if (txMappings.value == undefined) {
    return
  }
  let mapping = txMappings.value[idx]
  SetCointracking2BlockpitMapping(
      mapping.cointracking.value,
      mapping.blockpit.value,
  )
}

const setTimezone = async (newTz: string) => {
  store.timezone = newTz
  SetTimezone(newTz).catch((msg: any) => {
    console.log("failed setting timezone to backend " + msg)
  })
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

    <p class="text-h7 font-weight-bold mt-5">General</p>
    <TimezoneSelector
        :selected-timezone="store.timezone"
        @timezoneChanged="setTimezone"></TimezoneSelector>

    <p class="text-h7 font-weight-bold mt-5">Tx Type Mapping</p>
    <p class="text-subtitle-2">A mapping between CoinTracking and Blockpit transaction types</p>
    <v-table class="mt-4">
      <thead>
      <tr>
        <th class="text-left font-weight-bold">
          CoinTracking Transaction
        </th>
        <th class="text-left font-weight-bold">
          Blockpit Transaction
        </th>
      </tr>
      </thead>
      <tbody>
      <tr
          v-for="(mapping, index) in txMappings"
          :key="mapping.cointracking.value">
        <td>{{ mapping.cointracking.title }}</td>
        <td>
          <v-select
              density="compact"
              v-model="mapping.blockpit.value"
              :items="blockpitTxTypes"
              hide-details="auto"
              @update:model-value="blockpitTxTypeChanged(index)"
          ></v-select>
        </td>
      </tr>
      </tbody>
    </v-table>
  </TitledPanel>
</template>

<style scoped>

</style>