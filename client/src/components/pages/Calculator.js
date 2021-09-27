import React, {useState, useContext, useEffect} from 'react';
import bankContext from '../../context/bank/bankContext';

const About = () => {
    const context = useContext(bankContext)

    const {banks, GetBanks, GetBank, EraseBank, bank } = context

    useEffect(() => {

        GetBanks()
        if (bank === null) {
            setPayment("")
        }
        //eslint-disable-next-line
    }, [bank])

    console.log(banks)
    const [mortgage, setMortgage] = useState({
        init_loan:"",
        down_payment: ""
    })

    const [payment, setPayment] = useState("")

    const {init_loan, down_payment} = mortgage

    const onChange = (e) => setMortgage({
        ...mortgage,
        [e.target.name]: e.target.value
        
    })
    const onSelect = (e) => {
        if (e.target.value !== true) {
            GetBank(e.target.value)
            
        } else {
            EraseBank()
        }
        
        console.log(e.target.value)
    }

    const onSubmit = (e) => {
        e.preventDefault()

        setPayment(12344)
    }
    

    return (
        <div className="grid-2">
           <div>
           <form onSubmit={onSubmit}>
                <select name="bank" onChange={onSelect}>
                    <option value>Please Select Bank</option>
                    {banks === null ? (
                        <option disabled>No Records Found</option>
                    ) : banks.map((b) => <option value={b._id} key={b._id}>{b.name}</option>)
                    }
                </select>
                { (bank !== null) ?
                    <div>
                        <input type="number" 
                        placeholder="Enter Initial Loan" 
                        max={bank.max_loan}
                        name="init_loan" value={init_loan}
                        onChange={onChange}
                            required/>

                        <input type="number" 
                        placeholder="Enter Down Payment" 
                        max={bank.min_down}
                        name="down_payment" value={down_payment}
                        onChange={onChange}
                        required/>

                        <input type="submit" value="Mortgage Calculate" className="btn btn-block btn-primary" />
                    </div> : ""
                }
               
           </form>
                <div>
                    
                        <div>
                            <table className="table table-responsive">
                                <tr>
                                    <th>Month</th>
                                    <th>Total Paymemt</th>
                                    <th>Interest Paymemnt</th>
                                    <th>Loan Balance</th>
                                    <th>Equity</th>
                                </tr>
                            </table>
                        </div>
                    
                </div>
           </div>
            
        </div>
    )
}

export default About;