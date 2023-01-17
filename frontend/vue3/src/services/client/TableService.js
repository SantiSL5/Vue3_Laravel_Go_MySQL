import Api from '@/services/Api'
import secret from '../../Secret'

export default {
    getAllTables(data) {
        return Api(`${secret.GO_APP_URL}`).get(`table/`, { params: { offset: data.offset, limit: data.limit, capacity: data.capacity } })
    },
    getTableById(id) {
        return Api(`${secret.GO_APP_URL}`).get(`table/${id}`)
    }
}