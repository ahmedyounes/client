// @flow
import * as React from 'react'
import {WaitingButton} from '../../../common-adapters'
import * as Styles from '../../../styles'

type Props = {|
  following?: boolean,
  waitingKey: string,
  style?: Object,
  onFollow?: () => void,
  onUnfollow?: () => void,
|}

type State = {mouseOver: boolean}

class FollowButton extends React.Component<Props, State> {
  state: State

  constructor(props: Props) {
    super(props)

    this.state = {
      mouseOver: false,
    }
  }

  render() {
    const {following, onFollow, onUnfollow, style, waitingKey, ...otherProps} = this.props

    if (following) {
      return (
        <WaitingButton
          type="PrimaryGreenActive"
          label={this.state.mouseOver ? 'Unfollow' : 'Following'}
          onClick={onUnfollow}
          onMouseEnter={Styles.isMobile ? null : () => this.setState({mouseOver: true})}
          onMouseLeave={Styles.isMobile ? null : () => this.setState({mouseOver: false})}
          waitingKey={waitingKey}
          style={{...styleButton, ...style}}
          {...otherProps}
        />
      )
    } else {
      return (
        <WaitingButton
          type="PrimaryGreen"
          label="Follow"
          onClick={onFollow}
          waitingKey={waitingKey}
          style={{...styleButton, ...style}}
          {...otherProps}
        />
      )
    }
  }
}

const styleButton = Styles.platformStyles({
  isElectron: {
    width: 125,
  },
})

export default FollowButton
