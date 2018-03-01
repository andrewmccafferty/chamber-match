import { 
  ADD_USER,
} from '../../modules/constants'

const initialState = {
  users: []
}

const addUser = (state = initialState, action) => {
  switch (action.type) {
    case ADD_USER:
      return {
        ...state,
        users: [
          ...state.users,
          { 
            id: action.id,
            name: action.name,
            instrument: action.instrument,
            location: action.location
          }
        ]
      }
    default:
      return state
  }
}

export {
  addUser
}