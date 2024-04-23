<script setup lang="ts">

import {useRouter} from "vue-router";
import {inject, onMounted, ref} from "vue";
import {common} from "@wails/go/models";
import {useSettingsStore} from "@/stores/settings_store";
import TitledPanel from "@/components/TitledPanel.vue";
import TimezoneSelector from "@/components/TimezoneSelector.vue";
import InfoButton from "@/components/InfoButton.vue";
import {useSnackbarStore} from "@/stores/snackbar_store";
import {WailsApi} from "@/wails/wails_api";
import {wailsClientInjKey} from "@/injection_keys";

const router = useRouter()
const settingsStore = useSettingsStore()
const snackStore = useSnackbarStore()

const txMappings = ref<Array<common.Ct2BpTxMapping>>()
let blockpitTxTypes = ref(Array<common.TxDisplayName>())

const wailsClient: WailsApi = inject<WailsApi>(wailsClientInjKey) as WailsApi

const navigateBack = () => {
  router.back()
}

onMounted(() => {
  wailsClient.TxTypeMappings().then((mappings) => {
    txMappings.value = mappings
  }).catch((reason) => {
    snackStore.showError(`failed getting tx type mappings between cointracking and blockpit: ${reason}`)
  })

  wailsClient.BlockpitTxTypes().then((displayNames) => {
    blockpitTxTypes.value = displayNames
  }).catch((reason) => {
    snackStore.showError(`failed getting blockpit tx types: ${reason}`)
  })

  wailsClient.Timezone().then((timezone: string) => {
    settingsStore.timezone = timezone
  })
})

const blockpitTxTypeChanged = (idx: number) => {
  if (txMappings.value == undefined) {
    return
  }
  let mapping = txMappings.value[idx]
  wailsClient.SetCointracking2BlockpitMapping(
      mapping.cointracking.value,
      mapping.blockpit.value,
  )
}

const setTimezone = async (newTz: string) => {
  settingsStore.timezone = newTz
  wailsClient.SetTimezone(newTz).catch((msg: any) => {
    snackStore.showError("failed setting timezone to backend " + msg)
  })
}

const onSwapHandlingChanged = (newSwapHandling: string | null) => {
  if (!newSwapHandling) {
    return
  }
  wailsClient.SetSwapHandling(newSwapHandling).catch((msg: any) => {
    snackStore.showError(`failed changing swap handling: ${msg}`)
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
        ref="tzSelector"
        :selected-timezone="settingsStore.timezone"
        @timezoneChanged="setTimezone"></TimezoneSelector>

    <p class="text-h7 font-weight-bold mt-5">Swap Transaction Handling</p>
    <p class="text-subtitle-2">How to handle CoinTracking Swap transactions for Blockpit export.</p>
    <v-radio-group
        v-model="settingsStore.swapHandling"
        @update:model-value="onSwapHandlingChanged">
      <v-radio value="swap_non_taxable">
        <template #label>
          <v-label>Non Taxable</v-label>
          <InfoButton>
            A CoinTracking swap transaction will be split up into a "Non-Taxable Out" and a "Non-Taxable In" transaction
            for Blockpit.
          </InfoButton>
        </template>
      </v-radio>
      <v-radio value="swap_to_trade">
        <template #label>
          <v-label>Trade</v-label>
          <InfoButton>
            A CoinTracking swap transaction will be converted into a "Trade" transaction for Blockpit which is a taxable
            event.
          </InfoButton>
        </template>
      </v-radio>
    </v-radio-group>

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