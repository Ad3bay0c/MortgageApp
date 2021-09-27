import React, { useState, useContext, useEffect } from 'react';
import BankContext from '../../context/bank/bankContext';
// import ContactItem from './ContactItem';

const BankForm = () => {
    const bankContext = useContext(BankContext)

    const { CreateBank, updateBank, current, clearCurrent } = bankContext

    const [bank, setBank] = useState({
        name: "",
        interest: "",
        loan_term: "",
        max_loan: "",
        min_down: ""
    })

    useEffect(() => {
        if (current !== null) {
            setBank(current)
        } else {
            setBank({
                name: "",
                interest: "",
                loan_term: "",
                max_loan: "",
                min_down: ""
            })
        }
    }, [current])

    const { name, interest, loan_term, max_loan, min_down } = bank

    const submitForm = e => {
        e.preventDefault();
        bank.interest = Number(bank.interest)
        bank.loan_term = Number(bank.loan_term)
        bank.max_loan = Number(bank.max_loan)
        bank.min_down = Number(bank.min_down)

        CreateBank(bank)
        setBank({
            name: "",
            interest: "",
            loan_term: "",
            max_loan: "",
            min_down: ""
        })
    }
    const updateForm = e => {
        e.preventDefault();
        bank.interest = Number(bank.interest)
        bank.loan_term = Number(bank.loan_term)
        bank.max_loan = Number(bank.max_loan)
        bank.min_down = Number(bank.min_down)

        updateBank(bank, current._id)

        setBank({
            name: "",
            interest: "",
            loan_term: "",
            max_loan: "",
            min_down: ""
        })
    }

    const onChange = (e) => setBank({
        ...bank,
        [e.target.name]: e.target.value
    })
    return (
        <form onSubmit={current === null ? submitForm : updateForm}>
            <h2 className="text-Primary">{current === null ? "Create New Bank" : "Update Bank"}</h2>
            <input name="name" type="text" placeholder="Enter Bank Name" value={name} onChange={onChange} required />

            <input name="interest" type="number" placeholder="Enter Loan Interest" value={interest} onChange={onChange} required />

            <input name="loan_term" type="number" placeholder="Enter Loan Term in Month" value={loan_term} onChange={onChange} required />

            <input name="max_loan" type="number" placeholder="Enter Maximum Loan" value={max_loan} onChange={onChange} required />

            <input name="min_down" type="number" placeholder="Enter Minimum Down Payment" value={min_down} onChange={onChange} required />

            <input type="submit" className={`btn btn-${current === null ? "primary" : "dark"} btn-block`} value={current === null ? "Create Bank" : "Update Bank"} />

            {current && <div>
                <button className="btn btn-light btn-block"
                    onClick={clearCurrent}>
                    Clear
                </button>
            </div>}

        </form>
    );
}

export default BankForm;