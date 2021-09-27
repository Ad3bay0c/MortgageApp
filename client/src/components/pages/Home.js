import React, { useEffect, useContext } from 'react';
import ContactForm from '../banks/BankForm';
import Banks from '../banks/banks';
import bankContext from '../../context/bank/bankContext';
import axios from 'axios'
const Home = () => {
    const context = useContext(bankContext)

    const { GetBanks } = context

    useEffect(() => {

        GetBanks()
        //eslint-disable-next-line
    }, [])

    return (
        <div className="grid-2">
            <div>
                <ContactForm />
            </div>
            <div>
                <Banks />
            </div>
        </div>
    )
}

export default Home;