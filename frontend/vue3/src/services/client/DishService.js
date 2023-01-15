import Api from '@/services/Api'
import secret from '../../Secret'

export default {
    getAllDishes() {
        return Api(`${secret.GO_APP_URL}`).get(`dish/`)
    },
    getDishById(id) {
        return Api(`${secret.GO_APP_URL}`).get(`dish/${id}`)
    }
}
