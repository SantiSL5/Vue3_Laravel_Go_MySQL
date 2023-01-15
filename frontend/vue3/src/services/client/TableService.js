import Api from '@/services/Api'
import secret from '../../Secret'

export default {
    getAllTables() {
        return Api(`${secret.GO_APP_URL}`).get(`table`)
    },
    getTableById(id) {
        return Api(`${secret.GO_APP_URL}`).get(`table/${id}`)
    }
}