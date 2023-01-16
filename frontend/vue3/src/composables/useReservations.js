import { ref } from "vue";
import ReservationService from '../services/admin/ReservationService'

export const useReservationsToggle = (data, id) => {
    const reservations = ref([])
    ReservationService.updateReservation(data, id)
        .then(res => reservations.value = res.data)
        .catch(error => console.error(error))
    return reservations;
};