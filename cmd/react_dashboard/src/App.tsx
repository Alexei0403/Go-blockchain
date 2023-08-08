import { fetchBlockchainData } from "api/miner";
import AppFooter from "components/layout/AppFooter";
import AppHeader from "components/layout/AppHeader";
import AppInfo from "components/layout/AppInfo";
import Background from "components/layout/Background";
import BlockDiv from "components/BlockDiv";
import Loader from "components/shared/Loader";
import Notification from "components/shared/Notification";
import React, { useEffect, useState } from "react";
import styled from "styled-components";
import Wallet from "components/wallet/Wallet";
import UtilReducer from "store/UtilReducer";

const AppWrapper = styled.div`
  margin: 0;
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  overflow: auto;
`;

const ContentContainer = styled.div`
  display: flex;
  flex-direction: column;
  align-items: center;
`;

const WalletWrapperContainer = styled.div`
  display: flex;
  flex-wrap: wrap;
  justify-content: space-between;

  @media (max-width: 850px) {
    justify-content: center;
  }
`;

function App() {
  const [isLoading, setIsLoading] = useState(true);
  const [isError, setIsError] = useState({ message: "" });
  const [blockchain, setBlockchain] = useState<Block[]>([]);

  const [utilState, dispatchUtil] = React.useReducer(UtilReducer, {
    isActive: false,
    type: "info",
    message: "",
  });

  function fetchchainData() {
    dispatchUtil({
      type: "ON",
      payload: {
        type: "info",
        message: "Fetching blockchain data...",
      },
    });

    return fetchBlockchainData()
      .then((blocks) => {
        setBlockchain(blocks);
        dispatchUtil({
          type: "OFF",
          payload: null,
        });
      })
      .catch((error) => {
        dispatchUtil({
          type: "ON",
          payload: {
            type: "error",
            message: error.message,
          },
        });
      });
  }

  useEffect(() => {
    // Fetch blockchain data immediately on component mount
    fetchchainData();
    // Fetch blockchain data every 5 seconds
    const intervalId = setInterval(() => {
      fetchchainData();
    }, 5000);
    // Clear interval on component unmount
    return clearInterval(intervalId);
  }, []);

  return (
    <AppWrapper>
      <Background />
      <AppHeader title="Go Blockchain" />
      <ContentContainer className="App">
        <AppInfo />

        <WalletWrapperContainer>
          <Wallet type="Miner" />
          <Wallet type="User" />
        </WalletWrapperContainer>

        {utilState.isActive && (
          <Notification
            type={utilState.type}
            message={utilState.message}
            underDevelopment={true}
            insideContainer={false}
          />
        )}

        {!utilState.isActive &&
          blockchain.map((block, index) => (
            <React.Fragment key={index}>
              <Loader height={100} />
              <BlockDiv block={block} />
            </React.Fragment>
          ))}
      </ContentContainer>

      <AppFooter
        githubUrl="https://github.com/elarsaks"
        linkedinUrl="https://www.linkedin.com/in/elarsaks/"
        websiteUrl="https://saks.digital/"
      />
    </AppWrapper>
  );
}

export default App;
