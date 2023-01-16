import Api from '@/services/Api'
import secret from '../../Secret'

export default {
    getAllReservations() {
        return Api(`${secret.GO_APP_URL}`).get(`reserve/`)
    },
    getReservationById(id) {
        return Api(`${secret.GO_APP_URL}`).get(`reserve/${id}`)
    },
    createReservation(data) {
        return Api(`${secret.GO_APP_URL}`).post('reserve/', data)
    },
    updateReservation(data, id) {
        return Api(`${secret.GO_APP_URL}`).post(`reserve/${id}?_method=PUT`, data)
    },
    changeStatusReservation(data, id) {
        return Api(`${secret.GO_APP_URL}`).put(`reserve/changeStatus/${id}`, data)
    },
    deleteReservationById(id) {
        return Api(`${secret.GO_APP_URL}`).delete(`reserve/${id}`)
    }
}
