import Api from '@/services/Api'
import secret from '../../Secret'

export default {
    getAllDishTypes(offset, limit) {
        return Api(`${secret.GO_APP_URL}`).get(`dishtype/`, { params: { offset: offset, limit: limit } })
    },
    getDishTypeById(id) {
        return Api(`${secret.GO_APP_URL}`).get(`dishtype/${id}`)
    }
}
