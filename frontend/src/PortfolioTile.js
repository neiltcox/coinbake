import React from "react";
import InlineLayout from "./InlineLayout";
import Stack from './Stack';
import ContentBox from "./ContentBox";
import exchangeName from './exchangeName';
import StyledText from "./StyledText";
import ConnectionStatus from "./ConnectionStatus";
import CurrencyFormat from 'react-currency-format';

export default function PortfolioTile(props) {
	return (
		<div className='portfolio_tile'>
			<ContentBox link={"/portfolio/"+props.portfolio.ID}>
				<Stack gap_size='shim'>
					<StyledText styling='title'>{props.portfolio.Name}</StyledText>
					<InlineLayout gap_size='minor'>
						<StyledText styling='standard'>{exchangeName[props.portfolio.ExchangeIdentifier]}</StyledText>
						<ConnectionStatus sentiment={props.portfolio.Connected ? 'positive' : 'negative'} />
					</InlineLayout>
					<StyledText><CurrencyFormat value={props.portfolio.TotalValuation} displayType='text' prefix='$' thousandSeparator={true}/></StyledText>
				</Stack>
			</ContentBox>
		</div>
	);
}