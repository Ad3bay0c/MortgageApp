import {
    CLEAR_CURRENT,
    CREATE_CONTACT,
    DELETE_BANK,
    ERASE_BANK,
    GET_BANK,
    GET_BANKS,
    REMOVE_ALERT,
    SET_ALERT,
    SET_CURRENT,
    UPDATE_BANK
} from '../type'

export default (state, action) => {
    switch (action.type) {
        case CREATE_CONTACT:
            return {
                ...state,
                banks: [action.payload, ...state.banks]
            }
        case SET_CURRENT:
            return {
                ...state,
                current: action.payload
            }
        case UPDATE_BANK:
            return {
                ...state,
                banks: state.banks.map((bank) => (bank._id === action.payload.id
                    ? action.payload.bank
                    : bank
                )),
                current: null
            }
        case DELETE_BANK:
            return {
                ...state,
                banks: state.banks.filter(
                    (bank) => bank._id !== action.payload
                    )
            }
        case CLEAR_CURRENT:
            return {
                ...state,
                current: null
            }
        case ERASE_BANK:
            return {
                ...state,
                bank: null
            }
        case GET_BANKS:
            return {
                ...state,
                banks: action.payload
            }
        case GET_BANK:
            return {
                ...state,
                bank: action.payload
            }
        case SET_ALERT:
            return {
                ...state,
                alert: action.payload
            }
        case REMOVE_ALERT:
            return {
                ...state,
                alert: null
            }
    }
}