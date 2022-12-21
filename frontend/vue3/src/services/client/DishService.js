import Api from '@/services/Api'
import secret from '../../Secret'

export default {
    getAllDishes() {
        return Api(`${secret.GO_APP_URL}`).get(`dish/`)
    },
    getDishById(id) {
        return Api(`${secret.GO_APP_URL}`).get(`dish/${id}`)
    },
    createDish(data) {
        return Api(`${secret.GO_APP_URL}`).post('dish/', data)
    },
    updateDish(data, id) {
        return Api(`${secret.GO_APP_URL}`).post(`dish/${id}?_method=PUT`, data)
    },
    changeStatusDish(data, id) {
        return Api(`${secret.GO_APP_URL}`).put(`dish/changeStatus/${id}`, data)
    },
    deleteDishById(id) {
        return Api(`${secret.GO_APP_URL}`).delete(`dish/${id}`)
    }
}
