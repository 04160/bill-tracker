import React, { Component } from 'react';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';
import { getBills } from '../actions/actions';

class BillList extends Component {
  componentWillMount() {
    this.props.getBills();
  }

  componentWillReceiveProps(nextProps) {
    if (nextProps.newBill) {
      this.props.bills.unshift(nextProps);
    }
  }

  render () {
    console.log(this.props.bills);
    const billItems = this.props.bills.map(bill => (
      <div key={ bill.id }>
        <h2>{ bill.total }</h2>
        <p>{ bill.description }</p>
      </div>
    ));
    return (
      <div>
        <h1>Bills</h1>
        { billItems }
      </div>
    );
  }
}

BillList.propTypes = {
  getBills: PropTypes.func.isRequired,
  bills: PropTypes.array.isRequired,
  newBill: PropTypes.object
};

const mapStateToProps = state => ({
  bills: state.bills.items,
  newBill: state.bills.item
});

export default connect (mapStateToProps, { getBills }) (BillList);