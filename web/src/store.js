import { createStore, applyMiddleware, compose } from 'redux'
import { routerMiddleware } from 'react-router-redux'
import thunk from 'redux-thunk'
import createHistory from 'history/createBrowserHistory'
import rootReducer from './reducersIndex'

const history = createHistory()

const initialState = {}
const enhancers = []
const middleware = [
	thunk,
	routerMiddleware(history)
]

const devToolsExtension = window.devToolsExtension //taken out conditional to check NODE_ENV

if (typeof devToolsExtension === 'function') {
	enhancers.push(devToolsExtension())
}

const composedEnhancers = compose(
	applyMiddleware(...middleware),
	...enhancers
)

const store = createStore(
	rootReducer,
	initialState,
	composedEnhancers
)

export { 
	history,
	store as default
}