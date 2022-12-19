import Api from '@/services/Api'
import secret from '../../Secret'

export default {
    getAllDishTypes() {
        return Api(`${secret.GO_APP_URL}`).get(`dishtype/`)
    },
    getDishTypeById(id) {
        return Api(`${secret.GO_APP_URL}`).get(`dishtype/${id}`)
    },
    createDishType(data) {
        return Api(`${secret.GO_APP_URL}`).post('dishtype/', data)
    },
    updateDishType(data, id) {
        return Api(`${secret.GO_APP_URL}`).post(`dishtype/${id}?_method=PUT`, data)
    },
    changeStatusDishType(data, id) {
        return Api(`${secret.GO_APP_URL}`).put(`dishtype/changeStatus/${id}`, data)
    },
    deleteDishTypeById(id) {
        return Api(`${secret.GO_APP_URL}`).delete(`dishtype/${id}`)
    }
}
