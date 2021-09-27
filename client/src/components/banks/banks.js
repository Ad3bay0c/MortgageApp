import React, { Fragment, useContext } from 'react';
import bankContext from '../../context/bank/bankContext';
import BankItem from './bankItem';

const Banks = () => {
    const context = useContext(bankContext)

    const { banks } = context

    if (banks.length === 0) {

        return (
            <div>
                <h2 className="text-Primary">Available Banks</h2>
                <p>Empty List, Please Create One</p>
            </div>
        )
    }
    return (
        <Fragment>
            <h2 className="text-Primary">Available Banks</h2>
            {banks.map((bank) => <BankItem key={bank._id} Bank={bank} />)}
        </Fragment>
    )
}


export default Banks;