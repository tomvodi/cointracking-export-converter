import {defineStore} from "pinia";
import {computed, ref} from "vue";
import {common} from "../../wailsjs/go/models";


export const useSettingsStore = defineStore('settings', () => {
    const settingsLoaded = ref<boolean>(false)
    const timezone = ref<string>('')
    const allTimezones = ref<Array<common.TimezoneData>>([])
    const swapHandling = ref<string>("")

    const timezoneEmpty = computed(() => {
        return timezone.value.length == 0 && settingsLoaded.value
    })

    return {timezone, timezoneEmpty, allTimezones, settingsLoaded, swapHandling}
})