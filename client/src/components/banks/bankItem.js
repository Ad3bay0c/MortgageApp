import React, { useContext } from 'react'
// import {Link} from 'react-router-dom'

import PropTypes from 'prop-types';
import bankContext from '../../context/bank/bankContext';

const BankItem = ({ Bank }) => {
    const context = useContext(bankContext)

    const { setCurrent, deleteBank } = context

    const SetCurrent = () => {
        setCurrent(Bank)
    }
    const DeleteBank = () => {
        deleteBank(Bank._id)
    }
    const { name, interest, max_loan, loan_term } = Bank
    return (
        <div className='card bg-light'>
            <h3 className='text-primary text-left'>
                {name}
                <span
                    style={{ float: 'right', padding: '4px' }}
                    className='badge-success'>
                    {interest}%
                </span>
            </h3>
            <ul className="list">
                {max_loan && (<li>
                    <i className='fas fa-envelope-open'></i> Max_Loan: ${max_loan}
                </li>)}
                {loan_term && (<li>
                    <i className='fas fa-phone'></i> Loan Term: {loan_term} Month(s)
                </li>)}
            </ul>
            <p>
                <button className="btn btn-dark btn-sm"
                    onClick={SetCurrent}>Edit</button>
                <button className="btn btn-danger btn-sm"
                    onClick={DeleteBank}>Delete</button>
            </p>
        </div>
    )
}

BankItem.propTypes = {
    Bank: PropTypes.object.isRequired,
}
export default BankItem