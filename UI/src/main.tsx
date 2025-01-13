import ReactDOM from 'react-dom';
import { Provider } from 'react-redux';
import App from './App';
import './styles.css';
import { store } from './store/store.ts';

ReactDOM.render(
    <Provider store={store}>
        <App />
    </Provider>,
    document.getElementById('root')
);
