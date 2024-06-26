import {defineStore} from "pinia";
import {computed, ref} from "vue";
import {common} from "@wails/go/models";


export const useSettingsStore = defineStore('settings', () => {
    const settingsLoaded = ref<boolean>(false)
    const timezone = ref<string>('')
    const allTimezones = ref<Array<common.TimezoneData>>([])
    const swapHandling = ref<string>("")

    const timezoneEmpty = computed(() => {
        if (settingsLoaded.value) {
            return timezone.value === ''
        }
        return true
    })

    return {timezone, timezoneEmpty, allTimezones, settingsLoaded, swapHandling}
})