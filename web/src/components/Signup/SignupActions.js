import { 
  ADD_USER,
} from '../../modules/constants'

const addUser = (id, name, instrument, location) => {
  return dispatch => {
    dispatch({
      type: ADD_USER,
      id,
      name,
      instrument,
      location
    })
  }
}

export {
	addUser,
}