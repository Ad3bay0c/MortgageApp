import React, {useContext} from 'react'
import bankContext from '../../context/bank/bankContext'

const Alert = () => {
    const context = useContext(bankContext)

    const {alert} = context

    return (
        alert !== null && (
        <div className={`alert alert-${alert.color}`}>
        {/* {console.log(alert)} */}
           <i>{ alert.message }</i>
        </div>
        )
    );
}

export default Alert