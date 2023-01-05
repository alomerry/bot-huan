<template>
  <Card :style="{background: 'rgb(75 187 128 / 90%)',margin:'0 auto'}" class="danmaku">
    <Affix :offset-top="0">
      <Cell
          :style="{'text-align': 'center','backdrop-filter': 'saturate(150%) blur(20px)',transition: 'background-color var(--t-color),border-color var(--t-color)','border-radius':'5px'}">
        <Space align="center" class="danmaku-streamer">
          <Icon type="md-clock" :size="18"/>
          <Time :time="now" type="datetime" v-font="17"/>
          <Icon type="md-heart" :size="18"/>
          <CountUp :end="streamer.star" :duration="6" v-font="17"/>
          <Icon type="md-person" :size="18"/>
          <CountUp :end="streamer.watcher" :duration="6" v-font="17"/>

          <template v-if="streamer.status=='online'">
            <Badge status="success"/>
            <Tag color="success">online</Tag>
          </template>
          <template v-else>
            <Badge status="error"/>
            <Tag color="error">offline</Tag>
          </template>

        </Space>
      </Cell>
    </Affix>

    <Row>
      <Col span="24">
        <List item-layout="vertical" id="danmakuRef"
              :style="{'margin-bottom': '200px','overflow': 'hidden'}">
          <TransitionGroup name="list" tag="div">
            <template v-for="danmaku in danmakus" :key="danmaku.id">
              <Card class="danmaku-list-card" :bordered="false">
                <ListItem>
                  <Alert type="info">
                    <div v-line-clamp="4" v-font="15">
                      <template v-for="tag in danmaku.tags" :key="danmaku.tags">
                        <Avatar :src="tag" size="small" shape="square"/>&nbsp;
                      </template>
<!--                      <Avatar src="https://i.loli.net/2017/08/21/599a521472424.jpg"/>&nbsp;-->
                      <span v-color="danmaku.nameColor">
                        {{ danmaku.name }}
                      <Tag type="dot" :color="danmaku.nameColor" v-font="14" v-color="danmaku.nameColor+'!important;'">
                      </Tag>
                      </span>
                      <span v-if="danmaku.bitAnimatedIcons!=''">
                        <Avatar :src="danmaku.bitAnimatedIcons"/>&nbsp;
                      </span>
                      <span v-color="'#ffffff'" v-html="danmaku.msg"></span>
                    </div>
                  </Alert>
                </ListItem>
              </Card>
            </template>
          </TransitionGroup>
        </List>

      </Col>
    </Row>

    <div class="demo-typography-styles">
      <Text>View Design (default)</Text>
      <Text type="secondary">View Design (secondary)</Text>
      <Text type="success">View Design (success)</Text>
      <Text type="warning">View Design (warning)</Text>
      <Text type="danger">View Design (danger)</Text>
      <Text disabled>View Design (disabled)</Text>
      <Text mark>View Design (mark)</Text>
      <Text code>View Design (code)</Text>
      <Text keyboard>View Design (keyboard)</Text>
      <Text underline>View Design (underline)</Text>
      <Text delete>View Design (delete)</Text>
      <Text strong>View Design (strong)</Text>
      <Text italic>View Design (italic)</Text>
      <Link to="https://www.iviewui.com" target="_blank">
        View Design (Link)
      </Link>
    </div>
  </Card>

  <Button @click="renderFunc">测试热潮列车</Button>

  <GlobalFooter :links="footLinks" :copyright="copyright"/>
</template>

<script lang="ts">
import {defineComponent} from 'vue';
import tmi from 'tmi.js'

const client = new tmi.Client({
  options: {debug: false, messagesLogLevel: "info"},
  connection: {
    reconnect: true,
    secure: true
  },
  identity: {
    username: 'justinfan12345',
    password: 'oauth:kappa',
    // username: 'alomerry',
    // password: 'oauth:vt5e1ifhtuy8po0twhhf4apbfsk7cu'
  },
  channels: ["trey24k"]
});

type danmaku = {
  id: string | undefined
  clientNonce: any
  name: string
  nameColor: string
  userId: string | undefined
  username: string | undefined
  msg: any
  tags: never[] | (() => never[])
  highLightType: string
  mod: boolean
  subscriber: boolean

  bitAnimatedIcons: string
}

export default defineComponent({
  name: 'Danmaku',
  data() {
    return {
      hotTrain: { // 热潮列车
        status: "start"
      },
      danmakuRef: null as HTMLElement | null, // 弹幕姬 dom 对象
      footLinks: [
        {key: 'github', icon: 'logo-github', href: 'https://github.com/alomerry/bot-huan', blankTarget: true}
      ],
      copyright: 'Copyright © 2023 Authed by Alomerry Wu',

      msgId: 0,
      streamer: {
        status: "offline", // 是否在线
        star: 26, // 关注数
        watcher: 356, // 观看数
        name: '清欢丶', // 主播昵称
      },
      danmakus: [] as danmaku[], // 弹幕列表
      now: new Date(), // 当前时间
      timer: 0 // 秒级定时器
    }
  },
  created() {
    this.$Notice.config({
      top: 150,
      duration: 30
    });
  },
  mounted() {
    let _this = this;
    this.timer = setInterval(() => {
      _this.now = new Date();
      _this.getStarAndWatchCount()

    }, 1000)

    client.connect().catch(console.error);
    // 聊天
    client.on('message', (channel, userState, message, self) => {
      if (self) return;
      console.log(userState, message)

      if (!this.danmakuRef) {
        this.danmakuRef = document.getElementById("danmakuRef")
      }
      if (userState.emotes) {
        console.log('##### msg', userState, message);
      }

      // 移除数组 top 部分弹幕
      if (this.danmakus.length > 50) {
        this.danmakus = this.danmakus.slice(this.danmakus.length / 2, this.danmakus.length + 1)
      }
      // 自动滚动到最新弹幕
      // window.scrollTo(0, this.danmakuRef!.scrollHeight + 50);
      let msg = {
        id: userState.id,
        mod: false,
        subscriber: false,
        clientNonce: userState['client-nonce'],
        name: userState['display-name'] ? userState['display-name'] : "匿名", // 清欢丶
        nameColor: userState.color ? userState.color : this.randomColor16(), // 名字顏色
        userId: userState['user-id'],
        username: userState['username'],
        msg: userState.emotes ? this.formatEmotes(message, userState.emotes) : message,
        tags: this.getTags(userState),
        highLightType: this.getHighLightType(this.getDanmakuType("chat", userState), userState), // 高亮类型：info/warn/error/success

        bitAnimatedIcons: this.getBitAnimatedIcons(userState)
      }
      this.danmakus.push(msg)
    });
// tmi.js v1.9 更新版:
    //
    // [1] 'subscription' [誰]訂閱
    // [2] 'resub' [誰]重新訂閱
    // [3] 'subgift' [誰]被[誰]贈送訂閱
    // [4] 'anonsubgift' 匿名贈送給[誰]訂閱
    // [5] 'submysterygift' [誰]隨機贈送幾個訂閱(顯示增送的數量;此次和歷來)
    // [6] 'anonsubmysterygift' 匿名隨機贈送幾個訂閱(顯示增送的數量;此次)
    // [7] 'primepaidupgrade' [誰]使用Prime訂閱
    // [8] 'giftpaidupgrade' [已有訂閱的人]被[誰]贈送升級訂閱
    // [9] 'anongiftpaidupgrade' [已有訂閱的人]被匿名贈送升級訂閱
    client.on("subscription", (channel, username, method, message, userState) => {
      // Do your stuff.
      console.log('##### subscription', userState, message, username, method);
    });

    // 呼币
    client.on('cheer', (channel, userState, message) => {
      console.log('##### cheer', userState, message);
      if (userState['message-type'] === 'action') console.log('Cheer message included /me');
      if (userState['message-type'] === 'chat') console.log('Cheer message included /me');

      if (!this.danmakuRef) {
        this.danmakuRef = document.getElementById("danmakuRef")
      }
      let msg = {
        id: userState.id,
        mod: false,
        subscriber: false,
        emotes: null,
        clientNonce: userState['client-nonce'],
        name: userState['display-name'] ? userState['display-name'] : "匿名", // 清欢丶
        nameColor: userState.color ? userState.color : this.randomColor16(), // 名字顏色
        userId: userState['user-id'],
        username: userState['username'],
        msg: userState.emotes ? this.formatEmotes(message, userState.emotes) : message,
        tags: userState.badges ? () => {
          console.log(userState.badges);
          // let badges_key = Object.keys(userState.badges);
          // let badges_key_length = badges_key.length;
          return []
        } : [],
        highLightType: this.getHighLightType(this.getDanmakuType("chat", userState), userState), // 高亮类型：info/warn/error/success

        bitAnimatedIcons: this.getBitAnimatedIcons(userState)
      }
      this.danmakus.push(msg)
      // {
      // bits :"100"
      // color    :      "#FF0000"
      // display-name    :      "CardiacArrest__"
      // }
      // doc: https://github.com/tmijs/tmi.js/issues/151
      //
      // ex.
      // @badge-info=subscriber/18;badges=subscriber/3012,hype-train/2;bits=87;color=#00FF7F;display-name=yummy8765;emotes=;first-msg=0;flags=;id=40b8b1fb-bfda-4ce6-a991-208316afbd0f;mod=0;room-id=47281189;subscriber=1;tmi-sent-ts=1634832158483;turbo=0;user-id=53201088;user-type= :yummy8765!yummy8765@yummy8765.tmi.twitch.tv PRIVMSG #tetristhegrandmaster3 :doodlecheer87 老闆我要一碗大頂牛，再一盤沙西米，那個綜合牛雜也幫我切一盤，還有滷肉飯也一碗，謝謝
      //
      // bits=87;
      // PRIVMSG #tetristhegrandmaster3 :doodlecheer87 老闆我要一碗大頂牛，再一盤沙西米，那個綜合牛雜也幫我切一盤，還有滷肉飯也一碗，謝謝
      // doodlecheer87

    });
  },
  beforeDestroy() {
    if (this.timer) {
      clearInterval(this.timer);
    }

    if (client) {
      client.disconnect();
    }
  },
  methods: {
    getTags(userState) {
      let tags = [] as string[]
      if (userState.badges != null) {
        // admin
        if (('admin' in userState.badges) && userstate.badges.admin >= 1) {
          tags.push("https://static-cdn.jtvnw.net/badges/v1/9ef7e029-4cdf-4d4d-a0d5-e2b3fb2583fe/1")
        }

        // staff
        if (('staff' in userState.badges) && userState.badges.staff >= 1) {
          tags.push("https://static-cdn.jtvnw.net/badges/v1/d97c37bd-a6f5-4c38-8f57-4e4bef88af34/1")
        }

        // global_mod
        if (('global_mod' in userState.badges) && userState.badges.global_mod >= 1) {
          tags.push("https://static-cdn.jtvnw.net/badges/v1/9384c43e-4ce7-4e94-b2a1-b93656896eba/1")
        }

        // ambassador
        if (('ambassador' in userState.badges) && userState.badges.ambassador >= 1) {
          tags.push("https://static-cdn.jtvnw.net/badges/v1/2cbc339f-34f4-488a-ae51-efdf74f4e323/1")
        }

        // 主播
        if (('broadcaster' in userState.badges) && userState.badges.broadcaster >= 1) {
          tags.push("https://static-cdn.jtvnw.net/chat-badges/broadcaster.png")
        }

        // mod
        if (('moderator' in userState.badges) && userState.badges.moderator >= 1) {
          tags.push("https://static-cdn.jtvnw.net/chat-badges/mod.png")
        }

        // vip
        if (('vip' in userState.badges) && userState.badges.vip >= 1) {
          tags.push("https://static-cdn.jtvnw.net/badges/v1/b817aba4-fad8-49e2-b88a-7cc744dfa6ec/1")
        }

        // 订阅者
        if (('subscriber' in userState.badges) && userState.badges.subscriber >= 0) {
          tags.push("https://static-cdn.jtvnw.net/badges/v1/5d9f2208-5dd8-11e7-8513-2ff4adfae661/1")
        }

        // sub-gift-leader
        if (('sub-gift-leader' in userState.badges) && userState.badges['sub-gift-leader'] >= 1) {
          let subGiftLeaderUrl = '';
          if (userState.badges['sub-gift-leader'] == 1) {
            subGiftLeaderUrl = "https://static-cdn.jtvnw.net/badges/v1/21656088-7da2-4467-acd2-55220e1f45ad/1";
          }
          if (userState.badges['sub-gift-leader'] == 2) {
            subGiftLeaderUrl = "https://static-cdn.jtvnw.net/badges/v1/0d9fe96b-97b7-4215-b5f3-5328ebad271c/1";
          }
          if (userState.badges['sub-gift-leader'] == 3) {
            subGiftLeaderUrl = "https://static-cdn.jtvnw.net/badges/v1/4c6e4497-eed9-4dd3-ac64-e0599d0a63e5/1";
          }

          if (subGiftLeaderUrl.length > 0) {
            tags.push(subGiftLeaderUrl)
          }
        }

        // sub-gifter
        if (('sub-gifter' in userState.badges) && userState.badges['sub-gifter'] >= 1) {
          let subGifterUrl = 'https://static-cdn.jtvnw.net/badges/v1/f1d8486f-eb2e-4553-b44f-4d614617afc1/1';
          if (userState.badges['sub-gifter'] >= 1) {
            subGifterUrl = 'https://static-cdn.jtvnw.net/badges/v1/f1d8486f-eb2e-4553-b44f-4d614617afc1/1';
          }
          if (userState.badges['sub-gifter'] >= 5) {
            subGifterUrl = 'https://static-cdn.jtvnw.net/badges/v1/3e638e02-b765-4070-81bd-a73d1ae34965/1';
          }
          if (userState.badges['sub-gifter'] >= 10) {
            subGifterUrl = 'https://static-cdn.jtvnw.net/badges/v1/bffca343-9d7d-49b4-a1ca-90af2c6a1639/1';
          }
          if (userState.badges['sub-gifter'] >= 25) {
            subGifterUrl = 'https://static-cdn.jtvnw.net/badges/v1/17e09e26-2528-4a04-9c7f-8518348324d1/1';
          }
          if (userState.badges['sub-gifter'] >= 50) {
            subGifterUrl = 'https://static-cdn.jtvnw.net/badges/v1/47308ed4-c979-4f3f-ad20-35a8ab76d85d/1';
          }
          if (userState.badges['sub-gifter'] >= 100) {
            subGifterUrl = 'https://static-cdn.jtvnw.net/badges/v1/5056c366-7299-4b3c-a15a-a18573650bfb/1';
          }
          if (userState.badges['sub-gifter'] >= 250) {
            subGifterUrl = 'https://static-cdn.jtvnw.net/badges/v1/df25dded-df81-408e-a2d3-40d48f0d529f/1';
          }
          if (userState.badges['sub-gifter'] >= 500) {
            subGifterUrl = 'https://static-cdn.jtvnw.net/badges/v1/f440decb-7468-4bf9-8666-98ba74f6eab5/1';
          }
          if (userState.badges['sub-gifter'] >= 1000) {
            subGifterUrl = 'https://static-cdn.jtvnw.net/badges/v1/b8c76744-c7e9-44be-90d0-08840a8f6e39/1';
          }

          tags.push(subGifterUrl)
        }

        // bits-leader
        if (('bits-leader' in userState.badges) && userState.badges['bits-leader'] >= 1) {
          let bitsLeaderUrl = '';
          if (userState.badges['bits-leader'] == 1) {
            bitsLeaderUrl = "https://static-cdn.jtvnw.net/badges/v1/8bedf8c3-7a6d-4df2-b62f-791b96a5dd31/1";
          }
          if (userState.badges['bits-leader'] == 2) {
            bitsLeaderUrl = "https://static-cdn.jtvnw.net/badges/v1/f04baac7-9141-4456-a0e7-6301bcc34138/1";
          }
          if (userState.badges['bits-leader'] == 3) {
            bitsLeaderUrl = "https://static-cdn.jtvnw.net/badges/v1/f1d2aab6-b647-47af-965b-84909cf303aa/1";
          }

          if (bitsLeaderUrl.length > 0) {
            tags.push(bitsLeaderUrl)
          }
        }

        // bits
        if (('bits' in userState.badges) && userState.badges.bits >= 1) {
          let bitsBadgeUrl = 'https://static-cdn.jtvnw.net/badges/v1/73b5c3fb-24f9-4a82-a852-2f475b59411c/1';

          if (userState.badges.bits >= 1) {
            bitsBadgeUrl = 'https://static-cdn.jtvnw.net/badges/v1/73b5c3fb-24f9-4a82-a852-2f475b59411c/1';
          }
          if (userState.badges.bits >= 100) {
            bitsBadgeUrl = 'https://static-cdn.jtvnw.net/badges/v1/09d93036-e7ce-431c-9a9e-7044297133f2/1';
          }
          if (userState.badges.bits >= 1000) {
            bitsBadgeUrl = "https://static-cdn.jtvnw.net/badges/v1/0d85a29e-79ad-4c63-a285-3acd2c66f2ba/1"
          }
          if (userState.badges.bits >= 5000) {
            bitsBadgeUrl = "https://static-cdn.jtvnw.net/badges/v1/57cd97fc-3e9e-4c6d-9d41-60147137234e/1"
          }
          if (userState.badges.bits >= 10000) {
            bitsBadgeUrl = "https://static-cdn.jtvnw.net/badges/v1/68af213b-a771-4124-b6e3-9bb6d98aa732/1"
          }
          if (userState.badges.bits >= 25000) {
            bitsBadgeUrl = "https://static-cdn.jtvnw.net/badges/v1/64ca5920-c663-4bd8-bfb1-751b4caea2dd/1"
          }
          if (userState.badges.bits >= 50000) {
            bitsBadgeUrl = "https://static-cdn.jtvnw.net/badges/v1/62310ba7-9916-4235-9eba-40110d67f85d/1"
          }
          if (userState.badges.bits >= 75000) {
            bitsBadgeUrl = "https://static-cdn.jtvnw.net/badges/v1/ce491fa4-b24f-4f3b-b6ff-44b080202792/1"
          }
          if (userState.badges.bits >= 100000) {
            bitsBadgeUrl = "https://static-cdn.jtvnw.net/badges/v1/96f0540f-aa63-49e1-a8b3-259ece3bd098/1"
          }
          if (userState.badges.bits >= 200000) {
            bitsBadgeUrl = "https://static-cdn.jtvnw.net/badges/v1/4a0b90c4-e4ef-407f-84fe-36b14aebdbb6/1"
          }
          if (userState.badges.bits >= 300000) {
            bitsBadgeUrl = "https://static-cdn.jtvnw.net/badges/v1/ac13372d-2e94-41d1-ae11-ecd677f69bb6/1"
          }
          if (userState.badges.bits >= 400000) {
            bitsBadgeUrl = "https://static-cdn.jtvnw.net/badges/v1/a8f393af-76e6-4aa2-9dd0-7dcc1c34f036/1"
          }
          if (userState.badges.bits >= 500000) {
            bitsBadgeUrl = "https://static-cdn.jtvnw.net/badges/v1/f6932b57-6a6e-4062-a770-dfbd9f4302e5/1"
          }
          if (userState.badges.bits >= 600000) {
            bitsBadgeUrl = "https://static-cdn.jtvnw.net/badges/v1/4d908059-f91c-4aef-9acb-634434f4c32e/1"
          }
          if (userState.badges.bits >= 700000) {
            bitsBadgeUrl = "https://static-cdn.jtvnw.net/badges/v1/a1d2a824-f216-4b9f-9642-3de8ed370957/1"
          }
          if (userState.badges.bits >= 800000) {
            bitsBadgeUrl = "https://static-cdn.jtvnw.net/badges/v1/5ec2ee3e-5633-4c2a-8e77-77473fe409e6/1"
          }
          if (userState.badges.bits >= 900000) {
            bitsBadgeUrl = "https://static-cdn.jtvnw.net/badges/v1/088c58c6-7c38-45ba-8f73-63ef24189b84/1"
          }
          if (userState.badges.bits >= 1000000) {
            bitsBadgeUrl = "https://static-cdn.jtvnw.net/badges/v1/494d1c8e-c3b2-4d88-8528-baff57c9bd3f/1"
          }
          if (userState.badges.bits >= 1250000) {
            bitsBadgeUrl = "https://static-cdn.jtvnw.net/badges/v1/ce217209-4615-4bf8-81e3-57d06b8b9dc7/1"
          }
          if (userState.badges.bits >= 1500000) {
            bitsBadgeUrl = "https://static-cdn.jtvnw.net/badges/v1/c4eba5b4-17a7-40a1-a668-bc1972c1e24d/1"
          }
          if (userState.badges.bits >= 1750000) {
            bitsBadgeUrl = "https://static-cdn.jtvnw.net/badges/v1/183f1fd8-aaf4-450c-a413-e53f839f0f82/1"
          }
          if (userState.badges.bits >= 2000000) {
            bitsBadgeUrl = "https://static-cdn.jtvnw.net/badges/v1/7ea89c53-1a3b-45f9-9223-d97ae19089f2/1"
          }
          if (userState.badges.bits >= 2500000) {
            bitsBadgeUrl = "https://static-cdn.jtvnw.net/badges/v1/cf061daf-d571-4811-bcc2-c55c8792bc8f/1"
          }
          if (userState.badges.bits >= 3000000) {
            bitsBadgeUrl = "https://static-cdn.jtvnw.net/badges/v1/5671797f-5e9f-478c-a2b5-eb086c8928cf/1"
          }
          if (userState.badges.bits >= 3500000) {
            bitsBadgeUrl = "https://static-cdn.jtvnw.net/badges/v1/c3d218f5-1e45-419d-9c11-033a1ae54d3a/1"
          }
          if (userState.badges.bits >= 4000000) {
            bitsBadgeUrl = "https://static-cdn.jtvnw.net/badges/v1/79fe642a-87f3-40b1-892e-a341747b6e08/1"
          }
          if (userState.badges.bits >= 4500000) {
            bitsBadgeUrl = "https://static-cdn.jtvnw.net/badges/v1/736d4156-ac67-4256-a224-3e6e915436db/1"
          }
          if (userState.badges.bits >= 5000000) {
            bitsBadgeUrl = "https://static-cdn.jtvnw.net/badges/v1/3f085f85-8d15-4a03-a829-17fca7bf1bc2/1"
          }

          tags.push(bitsBadgeUrl)
        }

        // partner
        if (('partner' in userState.badges) && userState.badges['partner'] >= 1) {
          tags.push("https://static-cdn.jtvnw.net/badges/v1/d12a2e27-16f6-41d0-ab77-b780518f00a3/1")
        }

        // premium
        if (('premium' in userState.badges) && userState.badges['premium'] >= 1) {
          tags.push("https://static-cdn.jtvnw.net/badges/v1/bbbe0db0-a598-423e-86d0-f9fb98ca1933/1")
        }

        // turbo
        if (('turbo' in userState.badges) && userState.badges['turbo'] >= 1) {
          tags.push("https://static-cdn.jtvnw.net/badges/v1/bd444ec6-8f34-4bf9-91f4-af1e3428d80f/1")
        }

        // glhf
        if (('glhf-pledge' in userState.badges) && userState.badges['glhf-pledge'] >= 0) {
          tags.push("https://static-cdn.jtvnw.net/badges/v1/3158e758-3cb4-43c5-94b3-7639810451c5/1")
        }

        // glitchcon2020
        if (('glitchcon2020' in userState.badges) && userState.badges['glitchcon2020'] >= 1) {
          tags.push("https://static-cdn.jtvnw.net/badges/v1/1d4b03b9-51ea-42c9-8f29-698e3c85be3d/1")
        }

        // bits-charity
        if (('bits-charity' in userState.badges) && userState.badges['bits-charity'] >= 1) {
          tags.push("https://static-cdn.jtvnw.net/badges/v1/a539dc18-ae19-49b0-98c4-8391a594332b/1")
        }
      }

      return tags
    },
    // 获取主播关注和观看人数
    getStarAndWatchCount() {
      let _this = this
      return new Promise((resolve, reject) => {
        // axios.xxx
        let info, err
        if (err) {
          reject(err)
        }
        resolve(info)
      }).then(info => {
        _this.streamer.star++
        _this.streamer.watcher++
      })
    },
    renderFunc(fans) {
      fans = "清欢"
      this.$Notice.open({
        title: '感谢' + fans + '赠送的礼物',
        desc: 'The desc will hide when you set render.',
        render: h => {
          return h('span', [
            'This is created by ',
            h('a', 'render'),
            ' function'
          ])
        }
      });

      this.$Notice.open({
        title: 'xxx 点了屠夫',
        desc: 'The desc will hide when you set render.',
        render: h => {
          return h('span', [
            'This is created by ',
            h('a', 'render'),
            ' function'
          ])
        }
      });

      this.$Notice.open({
        title: 'xxx 点了一首歌',
        desc: 'The desc will hide when you set render.',
        render: h => {
          return h('span', [
            'This is created by ',
            h('a', 'render'),
            ' function'
          ])
        }
      });
    },
    // 转换文本中图片表情
    formatEmotes(text, emotes) {
      // 切割文本
      let splitText = text.split('');
      // 遍历 emotes map
      for (let emotesKey in emotes) {
        let locations = emotes[emotesKey];
        for (let location in locations) {
          let mote = locations[location];
          if (typeof mote == 'string') {
            mote = mote.split('-');
            mote = [parseInt(mote[0]), parseInt(mote[1])];
            let length = mote[1] - mote[0],
                empty = Array.apply(null, new Array(length + 1)).map(function () {
                  return ''
                });
            splitText = splitText.slice(0, mote[0]).concat(`<img class="emoticon" src="https://static-cdn.jtvnw.net/emoticons/v2/${emotesKey}/default/dark/1.0"/>`).concat(splitText.slice(mote[1] + 1, splitText.length));
          }
        }
      }
      console.log(splitText.join(''))
      return splitText.join('')
    },
    getDanmakuType(eventType, userState) {
      switch (eventType) {
        case "chat":
          if ('msg-id' in userState && userState['msg-id'] == "highlighted-message") {
            return "chat-highlight"
          }
          if ('msg-id' in userState && userState['msg-id'] == "announcement") {
            return "chat-announcement"
          }
          if (('bits' in userState) && (userState.bits > 0)) {
            return "chat-cheer"
          }
          if ('msg-id' in userState && userState['msg-id'] == 'midnightsquid') {
            return "chat-midnight"
          }

          return "chat-msg"
        case "":
      }
    },
    isChatMessage() {

    },
    getHighLightType(danmakuType, userState) {
      // 醒目标识
      if ('msg-id' in userState && userState['msg-id'] == "highlighted-message") {
        return "warn"
      }

      // 公告
      if ('msg-id' in userState && userState['msg-id'] == "announcement") {
        let announcementColorType = userState['msg-param-color'] || 'PRIMARY';
        // announcement_color_classname_arr['PRIMARY'] = 'announcement_message_primary';
        // announcement_color_classname_arr['BLUE'] = 'announcement_message_blue';
        // announcement_color_classname_arr['GREEN'] = 'announcement_message_green';
        // announcement_color_classname_arr['ORANGE'] = 'announcement_message_orange';
        // announcement_color_classname_arr['PURPLE'] = 'announcement_message_purple';
        return "??"
      }

      // cheer 呼币
      if (('bits' in userState) && (userState.bits > 0)) {
        let bitIcons = 'https://static-cdn.jtvnw.net/bits/dark/animated/gray/1';
        if (userState.bits >= 1)
          bitIcons = 'https://static-cdn.jtvnw.net/bits/dark/animated/gray/1';
        if (userState.bits >= 100)
          bitIcons = 'https://static-cdn.jtvnw.net/bits/dark/animated/purple/1';
        if (userState.bits >= 1000)
          bitIcons = 'https://static-cdn.jtvnw.net/bits/dark/animated/green/1';
        if (userState.bits >= 5000)
          bitIcons = 'https://static-cdn.jtvnw.net/bits/dark/animated/blue/1';
        if (userState.bits >= 10000)
          bitIcons = 'https://static-cdn.jtvnw.net/bits/dark/animated/red/1';
        if (userState.bits >= 100000)
          bitIcons = 'https://static-cdn.jtvnw.net/bits/dark/animated/gold/1';

        let bitName = 'bit';
        if (userState.bits > 1) {
          bitName = 'bits';
        }

        // let message_with_cheers = message_with_emotes.replace(/(?<=^|\s+)(([a-z0-9]{1,25}cheer|cheer|BibleThump|cheerwhal|Corgo|uni|ShowLove|Party|SeemsGood|Pride|Kappa|FrankerZ|HeyGuys|DansGame|CleGiggle|TriHard|Kreygasm|4Head|SwiftRage|NotLikeThis|FailFish|VoHiYo|PJSalt|MrDestructoid|bday|RIPCheer|Shamrock)([1-9]{1,1}[0-9]{0,4}))(?=$|\s+)/gi, (match,p1,p2,p3,offset,string)=>{
        //   //console.log(p1,p2,p3); // ex. cheer100, cheer, 100
        //   let new_p2 = p2.toLowerCase();
        //
        //   let new_p3 = 1;
        //   if(p3>=100) new_p3 = 100;
        //   if(p3>=1000) new_p3 = 1000;
        //   if(p3>=5000) new_p3 = 5000;
        //   if(p3>=10000) new_p3 = 10000;
        //   //if(p3>=100000) new_p3 = 100000; //贈點沒到這階,要總合才有
        //
        //   let cheer_value_span_class_name = `twitch_cheer${new_p3}`;
        //
        //   //let cheer_img_str_static = `<img src="https://d3aqoihi2n8ty8.cloudfront.net/actions/${new_p2}/dark/static/${new_p3}/1.png" alt="">${p3}`; //png靜態圖
        //   //let cheer_img_str_gif = `<img src="https://d3aqoihi2n8ty8.cloudfront.net/actions/${new_p2}/dark/animated/${new_p3}/1.gif" alt="">${p3}`; //gif動態圖
        //
        //   let cheer_img_str_gif = _createElement.img_string(`https://d3aqoihi2n8ty8.cloudfront.net/actions/${new_p2}/dark/animated/${new_p3}/1.gif`, '');
        //   let cheer_value_span = _createElement.span_string(p3, true, [cheer_value_span_class_name]);
        //
        //   cheer_img_str_gif = cheer_img_str_gif + cheer_value_span + ' ';
        //
        //   return cheer_img_str_gif;
        // });
      }

      return "info"
    },

    getBitAnimatedIcons(userState) {
      let bitAnimatedIcons = ""
      if (userState) {
        let bits = Number(userState.bits)
        if (bits >= 1) {
          bitAnimatedIcons = "https://static-cdn.jtvnw.net/bits/light/animated/gray/1"
        } else if (bits >= 100) {
          bitAnimatedIcons = "https://static-cdn.jtvnw.net/bits/light/animated/purple/1"
        } else if (bits >= 1000) {
          bitAnimatedIcons = "https://static-cdn.jtvnw.net/bits/light/animated/green/1"
        } else if (bits >= 5000) {
          bitAnimatedIcons = "https://static-cdn.jtvnw.net/bits/light/animated/blue/1"
        } else if (bits >= 10000) {
          bitAnimatedIcons = "https://static-cdn.jtvnw.net/bits/light/animated/red/1"
        } else if (bits >= 100000) {
          bitAnimatedIcons = "https://static-cdn.jtvnw.net/bits/light/animated/gold/1"
        }
      }
      return bitAnimatedIcons
    },
    // 十六进制颜色随机
    randomColor16() {
      let r = Math.floor(Math.random() * 256);
      let g = Math.floor(Math.random() * 256);
      let b = Math.floor(Math.random() * 256);
      let color = '#' + (Array(6).join('0') + (r.toString(16) + g.toString(16) + b.toString(16))).slice(-6);
      return color;
    }
  },
});
</script>

<style scoped>
.ivu-card-body {
  padding: 0px !important;
}


.list-move, /* 对移动中的元素应用的过渡 */
.list-enter-active,
.list-leave-active {
  transition: all 1.5s ease;
}

.list-enter-from,
.list-leave-to {
  opacity: 0;
  transform: translateX(30px);
}

/* 确保将离开的元素从布局流中删除
  以便能够正确地计算移动的动画。 */
.list-leave-active {
  position: absolute;
}

.layout {
  border: 1px solid #d7dde4;
  background: #f5f7f9;
  position: relative;
  border-radius: 4px;
  overflow: hidden;
}

.layout-logo {
  width: 100px;
  height: 30px;
  background: #5b6270;
  border-radius: 3px;
  float: left;
  position: relative;
  top: 15px;
  left: 20px;
}

.layout-nav {
  width: 420px;
  margin: 0 auto;
  margin-right: 20px;
}

.layout-footer-center {
  text-align: center;
}

.dev-run-preview .dev-run-preview-edit {
  display: none
}


.danmaku-list-card >>> div.ivu-card-body {
  padding: 0 0;
}

.danmaku-list-card >>> li.ivu-list-item {
  padding: 0 0;
}

.danmaku-list-card >>> .ivu-alert {
  margin: 0 0;
}

.danmaku-list-card >>> .ivu-alert {
  padding: 5px 5px;
}

.danmaku-list-card >>> .ivu-alert-warning {
  border: 0px solid #ffd77a;
}

.danmaku-list-card >>> .ivu-alert-info {
  border: 0px solid #abdcff;
  background-color: rgba(240, 250, 255, 0);
}

.danmaku-list-card >>> .ivu-card-body, .danmaku-list-card >>> .ivu-card {
  background: #fff0;
}

.danmaku-list-card {
  font-weight: 900;
}

#danmakuRef >>> .ivu-list-container, #danmakuRef >>> .ivu-card {
  background: #fff0;
}

#danmakuRef >>> .ivu-tag, .danmaku-list-card.ivu-tag {
  border: 0px solid #e8eaec !important;
  margin: 1px -5px 1px -10px !important;
}

#danmakuRef >>> .ivu-tag-dot, .danmaku-list-card >>> .ivu-tag-dot {
  background: #fff0 !important;
}

.danmaku-list-card >>> .ivu-tag-dot-inner {
  margin: 0 0 !important;
  padding: 0 0 !important;

}

#danmakuRef >>> .ivu-tag-text, .danmaku-list-card >>> .ivu-tag-text {

}

.danmaku-streamer {
  color: white !important;
}

.danmaku >>> .ivu-cell {
  padding: 7px 0 !important;
}
</style>
