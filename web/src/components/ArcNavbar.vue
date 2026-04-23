<script setup lang="ts">
import { computed, ref } from 'vue'
import { useRouter } from 'vue-router'
import { useDisplay, useTheme } from 'vuetify'

const menuOpen = ref(false)
const router = useRouter()
const theme = useTheme()
const { width, height } = useDisplay()

const arcRadius = computed(() =>
  Math.max(180, Math.round(Math.min(width.value, height.value) / 3)),
)

const routeCircles = [
  { icon: 'mdi-theme-light-dark', ariaLabel: '切換主題', action: 'toggle-theme' },
  { icon: 'mdi-home', to: '/', ariaLabel: '首頁' },
  { icon: 'mdi-information-outline', to: '/about', ariaLabel: 'about' },
  { icon: 'mdi-login', to: '/login', ariaLabel: '登入' },
  { icon: 'mdi-account-circle-outline', to: '/profile', ariaLabel: 'profile' },
]

const goTo = async (to: string) => {
  menuOpen.value = false
  await router.push(to)
}

const toggleTheme = () => {
  theme.global.name.value =
    theme.global.current.value.dark ? 'lightTheme' : 'darkTheme'
}

const onArcItemClick = async (item: { action?: string; to?: string }) => {
  if (item.action === 'toggle-theme') {
    toggleTheme()
    return
  }
  if (item.to) {
    await goTo(item.to)
  }
}

const arcButtons = computed(() => {
  const start = 18
  const end = 78
  const step = (end - start) / (routeCircles.length - 1)
  const innerRadius = arcRadius.value - 42

  return routeCircles.map((item, index) => {
    const degree = start + step * index
    const rad = (degree * Math.PI) / 180

    return {
      ...item,
      left: Math.cos(rad) * innerRadius,
      top: Math.sin(rad) * innerRadius,
      index,
    }
  })
})
</script>

<template>
  <v-tooltip text="開啟導覽選單" location="right">
    <template #activator="{ props }">
      <v-btn
        v-bind="props"
        class="menu-button"
        icon
        size="large"
        aria-label="open menu"
        @click="menuOpen = !menuOpen"
      >
        <v-icon icon="mdi-menu" />
      </v-btn>
    </template>
  </v-tooltip>

  <transition name="fade">
    <div v-if="menuOpen" class="arc-mask" @click="menuOpen = false">
      <div class="aurora-layer" />
      <svg
        class="arc-shape"
        :viewBox="`0 0 ${arcRadius} ${arcRadius}`"
        :style="{ width: `${arcRadius}px`, height: `${arcRadius}px` }"
        aria-hidden="true"
      >
        <defs>
          <linearGradient id="arcFillGradient" x1="0%" y1="0%" x2="100%" y2="100%">
            <stop offset="0%" stop-color="#93c5fd" stop-opacity="0.32" />
            <stop offset="65%" stop-color="#818cf8" stop-opacity="0.12" />
            <stop offset="100%" stop-color="#22d3ee" stop-opacity="0.03" />
          </linearGradient>
          <linearGradient id="arcStrokeGradient" x1="0%" y1="100%" x2="100%" y2="0%">
            <stop offset="0%" stop-color="#7dd3fc" stop-opacity="0.95" />
            <stop offset="55%" stop-color="#a78bfa" stop-opacity="0.9" />
            <stop offset="100%" stop-color="#f0abfc" stop-opacity="0.86" />
          </linearGradient>
        </defs>
        <path
          :d="`M0 0 L${arcRadius} 0 A ${arcRadius} ${arcRadius} 0 0 1 0 ${arcRadius} Z`"
          class="arc-fill"
        />
        <path
          :d="`M${arcRadius} 0 A ${arcRadius} ${arcRadius} 0 0 1 0 ${arcRadius}`"
          class="arc-stroke"
        />
      </svg>

      <v-tooltip
        v-for="item in arcButtons"
        :key="item.to ?? item.action"
        :text="item.ariaLabel"
        location="right"
      >
        <template #activator="{ props }">
          <v-btn
            v-bind="props"
            class="route-circle"
            size="small"
            icon
            :aria-label="item.ariaLabel"
            :style="{
              left: `${item.left}px`,
              top: `${item.top}px`,
              '--i': item.index,
            }"
            @click.stop="onArcItemClick(item)"
          >
            <v-icon :icon="item.icon" />
          </v-btn>
        </template>
      </v-tooltip>
    </div>
  </transition>
</template>

<style scoped>
.menu-button {
  position: fixed;
  top: 16px;
  left: 16px;
  z-index: 1300;
  border: 1px solid rgba(125, 211, 252, 0.72);
  background: linear-gradient(145deg, #60a5fa, #2563eb);
  color: #f8fafc;
  box-shadow:
    0 8px 20px rgba(37, 99, 235, 0.38),
    inset 0 0 0 1px rgba(255, 255, 255, 0.26);
}

.menu-button::before {
  display: none;
}

.arc-mask {
  position: fixed;
  inset: 0;
  z-index: 1150;
  background:
    radial-gradient(120% 90% at 8% 12%, rgba(56, 189, 248, 0.18), transparent 45%),
    radial-gradient(80% 65% at 24% 30%, rgba(129, 140, 248, 0.16), transparent 60%),
    rgba(2, 6, 23, 0.72);
  backdrop-filter: blur(4px) saturate(130%);
}

.arc-shape {
  position: absolute;
  top: 0;
  left: 0;
  filter: drop-shadow(0 0 22px rgba(96, 165, 250, 0.25));
}

.aurora-layer {
  position: absolute;
  top: 0;
  left: 0;
  width: min(55vw, 460px);
  height: min(55vw, 460px);
  border-radius: 0 0 100% 0;
  background: radial-gradient(
    100% 100% at 0% 0%,
    rgba(125, 211, 252, 0.24) 0%,
    rgba(129, 140, 248, 0.18) 36%,
    rgba(244, 114, 182, 0.1) 62%,
    transparent 100%
  );
  animation: aurora-pulse 2.8s ease-in-out infinite;
  pointer-events: none;
}

.arc-fill {
  fill: url(#arcFillGradient);
}

.arc-stroke {
  fill: none;
  stroke: url(#arcStrokeGradient);
  stroke-width: 3;
  stroke-linecap: round;
  filter: drop-shadow(0 0 8px rgba(125, 211, 252, 0.5));
}

.route-circle {
  position: absolute;
  z-index: 1200;
  transform: translate(-50%, -50%);
  min-width: 34px;
  background: linear-gradient(160deg, rgba(186, 230, 253, 0.95), rgba(167, 139, 250, 0.95));
  color: #1e1b4b;
  border: 1px solid rgba(224, 231, 255, 0.9);
  box-shadow:
    0 10px 24px rgba(76, 29, 149, 0.35),
    0 0 0 4px rgba(56, 189, 248, 0.12);
  animation:
    pop-in 420ms cubic-bezier(0.22, 1, 0.36, 1) both,
    floaty 2.8s ease-in-out infinite;
  animation-delay:
    calc(var(--i) * 70ms),
    calc(var(--i) * 180ms);
}

.route-circle:hover {
  transform: translate(-50%, -50%) scale(1.08);
}

.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.25s ease;
}

.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}

@keyframes pop-in {
  from {
    transform: translate(-50%, -50%) scale(0.2) rotate(-20deg);
    opacity: 0;
    filter: blur(5px);
  }
  to {
    transform: translate(-50%, -50%) scale(1);
    opacity: 1;
    filter: blur(0);
  }
}

@keyframes floaty {
  0%,
  100% {
    translate: 0 0;
  }
  50% {
    translate: 0 -4px;
  }
}

@keyframes aurora-pulse {
  0%,
  100% {
    opacity: 0.8;
    filter: blur(0);
  }
  50% {
    opacity: 1;
    filter: blur(2px);
  }
}
</style>
