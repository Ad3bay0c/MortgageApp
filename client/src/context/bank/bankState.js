import { useReducer } from 'react'
import bankContext from './bankContext'
import BankReducer from './bankReducer'
import axios from 'axios'
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
const BankState = props => {
    const initialState = {
        current: null,
        banks: [],
        bank: null,
        alert: null
    }

    const config = {
        headers: {
            "Content-Type": "application/json"
        }
    }
    const [state, dispatch] = useReducer(BankReducer, initialState)

    // Get All Banks
    const GetBanks = async () => {

        try {
            const resp = await axios.get("http://localhost:5000/")

            // console.log(resp.data.data)
            dispatch({
                type: GET_BANKS,
                payload: resp.data.data
            })
        } catch (error) {
            console.log(error)
        }
    }
    //Create Contact
    const CreateBank = async (bank) => {
        try {
            const config = {
                headers: {
                    "Content-Type": "application/json"
                }
            }
            const res = await axios.post("http://localhost:5000/create", bank, config)
            bank._id = res.data.data
            console.log(bank)
            dispatch({
                type: CREATE_CONTACT,
                payload: bank
            })

            SetAlert("Bank Created Successfully", "success")
        } catch (err) {
            console.log(err)
            // SetAlert("Error Creating Bank", "danger")
        }

    }

    //Set Current
    const setCurrent = (bank) => {
        dispatch({
            type: SET_CURRENT,
            payload: bank
        })
    }

    //Clear Current
    const clearCurrent = () => {
        dispatch({
            type: CLEAR_CURRENT
        })
    }
    //Update Contact
    const updateBank = async (bank, id) => {
        const config = {
            headers: {
                "Content-Type": "application/json"
            }
        }
        try {
            const res = await axios.put(`http://localhost:5000/update/${id}`, bank, config)
            bank._id = id
            dispatch({
                type: UPDATE_BANK,
                payload: { bank: bank, id: id },
            })

            SetAlert("Updated Successfully", "success")
        } catch (err) {
            console.log(err)
        }
    }

    // Ddelete Bank 
    const deleteBank = async (id) => {
        try {
            const res = await axios.delete(`http://localhost:5000/delete/${id}`, config)

            dispatch({
                type: DELETE_BANK,
                payload: res.data.data
            })
            SetAlert("Deleted Successfully", "dark")
        } catch (err) {
            console.log(err)
        }
    }
    //Set Alert
    const SetAlert = (message, color) => {
        dispatch({
            type: SET_ALERT,
            payload: { message, color }
        })

        setTimeout(() => {
            dispatch({ type: REMOVE_ALERT })
        }, 5000)
    }
    //Get A Bank
    const GetBank = async (id) => {
        try {
            const res = await axios.get(`http://localhost:5000/${id}`)
            dispatch({
                type: GET_BANK,
                payload: res.data
            })
        } catch (err) {
            console.log(err)
        }
    }

    //Erase Bank
    const EraseBank = () => {
        dispatch({
            type: ERASE_BANK
        })
    }
    //Remove Alert
    const RemoveAlert = () => dispatch({
        type: REMOVE_ALERT
    })
    return (
        <bankContext.Provider
            value={{
                banks: state.banks,
                bank: state.bank,
                alert: state.alert,
                current: state.current,
                CreateBank,
                SetAlert,
                RemoveAlert,
                GetBanks,
                setCurrent,
                clearCurrent,
                updateBank,
                deleteBank,
                GetBank,
                EraseBank
            }}
        >
            {props.children}
        </bankContext.Provider>
    )
}

export default BankState
