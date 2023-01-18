import Api from '@/services/Api'
import secret from '../../Secret'

export default {
    getAllTables(data) {
        if (data.category == "") {
            return Api(`${secret.GO_APP_URL}`).get(`reserve/tables`, { params: { offset: data.offset, limit: data.limit, capacity: data.capacity, date: data.date, turn: data.turn } })
        } else {
            return Api(`${secret.GO_APP_URL}`).get(`reserve/tables`, { params: { offset: data.offset, limit: data.limit, capacity: data.capacity, date: data.date, turn: data.turn, category: data.category } })
        }
    },
    getTableById(id) {
        return Api(`${secret.GO_APP_URL}`).get(`table/${id}`)
    }
}