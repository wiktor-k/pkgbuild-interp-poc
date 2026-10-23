# Maintainer: Jakub Klinkovský <lahwaacz at archlinux dot org>

pkgname=nsight-systems
pkgver=2026.3.2.476
pkgrel=1
pkgdesc="System-wide statistical sampling profiler with tracing features"
arch=(x86_64 aarch64)
url="https://developer.nvidia.com/nsight-systems"
license=(LicenseRef-NVIDIA-SLA)
depends=(
  dbus
  fontconfig
  glib2
  glibc
  libdrm
  libglvnd
  libx11
  libxcb
  libxext
  libxi
  libxkbcommon
  libxkbcommon-x11
  libxkbfile
  libxshmfence
  nspr
  nss
  wayland
  xcb-util-cursor
  xcb-util-image
  xcb-util-keysyms
  xcb-util-renderutil
  # there are bundled libs: boost, gcc-libs, icu, libssh, libz, openssl, qt6, zstd
)
optdepends=(
  'cuda: for GPU profiling on the local system'
)
options=(!debug !strip)
source_x86_64=(https://developer.download.nvidia.com/compute/cuda/redist/${pkgname/-/_}/linux-x86_64/${pkgname/-/_}-linux-x86_64-$pkgver-archive.tar.xz)
source_aarch64=(https://developer.download.nvidia.com/compute/cuda/redist/${pkgname/-/_}/linux-sbsa/${pkgname/-/_}-linux-sbsa-$pkgver-archive.tar.xz)
source=(nsight-systems.desktop)

b2sums=('3f4fdc522ede1f4efa3ba91d369bbd1d576b0541f711558c309a697d33a1c9125c1ce3d430acc4e4ba060fa112ccb89c0d7e6284473b566502ca838202b790cc')
b2sums_x86_64=('6ae70fe98a597c1c248266d1587930d3c2c8d25cc3efabc957d3dd63043438583cfb54d764f7d4b3dd8b4fcd0eef6053a499f1f372032b59229edbd64c04581f')
b2sums_aarch64=('5f4b40a49eb5fce3cf0fe7aeaff0d8f7509a63c1c85d7429c9721812defc78025b8ac4bfc877176ec69b5cf12c2d72f399ff3dbd6222ac3947cc2d6ec77644e5')

package() {
  cd ${pkgname/-/_}-linux-*-${pkgver}-archive

  install -vdm 755 "$pkgdir"/usr/lib/$pkgname/
  cp -arv bin host-linux-* target-linux-* "$pkgdir"/usr/lib/$pkgname/

  # create symlinks in /usr/bin/
  install -vdm 755 "$pkgdir"/usr/bin
  ln -s -t "$pkgdir"/usr/bin/ ../lib/$pkgname/bin/{nsys,nsys-ui}

  # install documentation
  install -vdm 755 "$pkgdir"/usr/share/doc/$pkgname/
  cp -arv documentation/* "$pkgdir"/usr/share/doc/$pkgname/

  # install license
  install -vDm 644 LICENSE -t "$pkgdir"/usr/share/licenses/$pkgname/

  # install desktop entry
  install -vDm 644 -t "$pkgdir"/usr/share/applications/ ../$pkgname.desktop
}
