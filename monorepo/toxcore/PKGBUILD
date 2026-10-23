# Maintainer: Jiachen Yang <farseerfc@gmail.com>
# Maintainer: Carl Smedstad <carsme@archlinux.org>
# Contributor: Felix Yan <felixonmars@archlinux.org>
# Contributor: Håvard Pettersson <mail@haavard.me>
# Contributor: naxuroqa <naxuroqa at gmail.com>
# Contributor: Boohbah <boohbah at gmail.com>
# Contributor: Kevin MacMartin <prurigro at gmail dot com>

pkgname=toxcore
_pkgname=c-toxcore
epoch=1
pkgver=0.2.23
pkgrel=1
pkgdesc='Secure, configuration-free, P2P Skype replacement backend'
arch=('x86_64')
url='https://tox.chat'
license=('GPL-3.0-or-later')
depends=(
  'glibc'
  'libconfig'
  'libsodium'
  'libvpx'
  'opus'
)
makedepends=(
  'cmake'
  'git'
  'gtest'
)
provides=("tox")
conflicts=("tox")
backup=('etc/tox-bootstrapd.conf')
source=(
  "git+https://github.com/TokTok/c-toxcore#tag=v$pkgver?signed"
  "git+https://github.com/TokTok/cmp.git"
  'toxcore.sysusers'
  'toxcore.tmpfiles'
)
b2sums=('d61e87d20ea0c173873d27a11a24fc64192857f30a6590bc9b0ff2647cfedc111bb978eec586fbb4f0b2fe540a169db0b65403db788cc771baa10070652f6cc7'
        'SKIP'
        '8faed598106644d161fd59a369833bd985efd2b1da2611fb346e33dc1a4bf1cadc9863d0da3bc3fca74636b302ce7b721f764499e2714df6af4383ebb6c76996'
        '6c48116f064d9da91ea17a1b08d659b8f0201697676e478e7d7d80a05566029ea415ab270a799747c98e8c7cff4b0e1cb5448e21942531df90c8cc60e7032c6b')
validpgpkeys=(
  '15D3B9A6B3951DF9854FCA93E786548AE0A0B56B' # RobinLindén<mail+gitlab@robinlinden.eu>
  'B7008FF12C079BF6E6EA19D1601A604B7E605776' # Robin Lindén <mail@robinlinden.eu>
  # GPG key fingerprints taken from https://github.com/qTox/qTox/blob/$_gitver/README.md
  'DA262CC93C0E1E525AD21C8596775D454B8EBF44' # sudden6 <sudden6@gmx.at>
  'BA7883E22F9D35945BA3376053137C3033F09008' # Zetok Zalbavar <zetok@openmailbox.org>
  '1157616BBD860C539926F8139591A163FF9BE04C' # - antis81
  '31039166FA902CA50D05D6085AF9F2E29107C727' # - Diadlo
  'C7A2552D0B250F983827742C133203A3AC399151' # - initramfs
  '2880C860D95C909D3DA45C687E086DD661263264' # - tux3
  '94CB3CD51379B7734C5617D82DEB4706E996351E' # Green-Sky (Dev keys) <Green-Sky@users.noreply.github.com>
  'A90E2D9711FDD3A1DD68882C3855DBA2D74403C9' # Iphigenia Df <iphydf@users.noreply.github.com>
)

prepare() {
  cd $_pkgname
  git submodule init
  git config submodule.third_party/cmp.url "$srcdir/cmp"
  git -c protocol.file.allow=always submodule update \
    third_party/cmp

  sed -i "s|/usr/local|/usr|" other/bootstrap_daemon/tox-bootstrapd.service
  sed -i "/Rpath/d;/RPATH/d" CMakeLists.txt
}

build() {
  cd $_pkgname
  cmake -S . -B build \
    -DCMAKE_INSTALL_PREFIX=/usr \
    -DCMAKE_BUILD_TYPE=None \
    -Wno-dev \
    -DDHT_BOOTSTRAP=ON \
    -DBOOTSTRAP_DAEMON=ON \
    -DMUST_BUILD_TOXAV=ON \
    -DBUILD_MISC_TESTS=ON
  cmake --build build
}

check() {
  cd $_pkgname
  cmake --build build --target test
}

package() {
  cd $_pkgname
  DESTDIR="$pkgdir" cmake --install build
  install -vDm644 -t "$pkgdir/etc" other/bootstrap_daemon/tox-bootstrapd.conf
  install -vDm644 -t "$pkgdir/usr/lib/systemd/system" \
    other/bootstrap_daemon/tox-bootstrapd.service
  install -vDm644 ../toxcore.sysusers "$pkgdir/usr/lib/sysusers.d/toxcore.conf"
  install -vDm644 ../toxcore.tmpfiles "$pkgdir/usr/lib/tmpfiles.d/toxcore.conf"
}
