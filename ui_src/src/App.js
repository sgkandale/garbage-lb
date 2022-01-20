import Navigation from "./components/navigation";
import React from 'react';
import { Provider } from 'react-redux';
import { store } from './components/globalState'
import { ThemeProvider } from '@mui/material/styles'
import { theme } from './components/theme'

export default function App() {
	return <div className="App">
		<Provider store={store}>
			<ThemeProvider theme={theme}>
				<Navigation />
			</ThemeProvider>
		</Provider>
	</div>
}