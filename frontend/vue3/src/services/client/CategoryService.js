import Api from '@/services/Api'
import secret from '../../Secret'

export default {
    getAllCategories() {
        return Api(`${secret.GO_APP_URL}`).get(`category`)
    },
    getCategoryById(id) {
        return Api(`${secret.GO_APP_URL}`).get(`category/${id}`)
    }
}
