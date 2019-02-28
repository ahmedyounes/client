// @flow
import CoinFlip from '.'
import * as RPCChatTypes from '../../../../constants/types/rpc-chat-gen'
import {namedConnect} from '../../../../util/container'

type OwnProps = {|flipGameID: string|}

const participantStatuses = [
  RPCChatTypes.chatUiUICoinFlipPhase.commitment,
  RPCChatTypes.chatUiUICoinFlipPhase.reveals,
  RPCChatTypes.chatUiUICoinFlipPhase.complete,
]

const noParticipants = []

const mapStateToProps = (state, {flipGameID}: OwnProps) => {
  const status = state.chat2.getIn(['flipStatusMap', flipGameID])
  return !status
    ? {
        commitmentVis: '',
        participants: noParticipants,
        progressText: '',
        resultText: '',
        revealVis: '',
        showParticipants: false,
      }
    : {
        commitmentVis: status.commitmentVisualization,
        errorInfo: status.errorInfo,
        isError: status.phase === RPCChatTypes.chatUiUICoinFlipPhase.error,
        participants: status.participants || [],
        progressText: status.progressText,
        resultInfo: status.resultInfo,
        resultText: status.resultText,
        revealVis: status.revealVisualization,
        showParticipants: participantStatuses.indexOf(status.phase) >= 0,
        status,
      }
}

export default namedConnect<OwnProps, _, _, _, _>(
  mapStateToProps,
  (d, o) => ({}),
  (s, d, o) => ({...s}),
  'CoinFlip'
)(CoinFlip)
