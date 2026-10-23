# Maintainer: Antonio Rojas <arojas@archlinux.org>
# Contributor: Lukas Jirkovsky <l.jirkovsky@gmail.com>

pkgname=krusader
pkgver=2.9.0
pkgrel=2
pkgdesc='Advanced twin panel (commander style) file manager'
arch=(x86_64)
url='https://krusader.org/'
license=(GPL-2.0-or-later)
depends=(acl
         glibc
         karchive
         kbookmarks
         kcodecs
         kcolorscheme
         kcompletion
         kconfig
         kconfigwidgets
         kcoreaddons
         kcrash
         kglobalaccel
         kguiaddons
         ki18n
         kiconthemes
         kio
         kitemviews
         kjobwidgets
         kparts
         kservice
         kstatusnotifieritem
         ktextwidgets
         kwallet
         kwidgetsaddons
         kwindowsystem
         kxmlgui
         libstdc++
         qt6-5compat
         qt6-base
         solid
         zlib)
makedepends=(extra-cmake-modules
             kdoctools)
optdepends=('xz: LZMA and XZ archive support'
            'unzip: ZIP decompression support'
            'zip: ZIP archive support'
            'lhasa: LHA archive support'
            'cpio: cpio archive support'
            'unrar: RAR decompression support'
            'arj: ARJ archive support'
            'unarj: ARJ decompression support'
            'unace: ACE decompression support'
            '7zip: 7z archive support'
            # additional apps
            'keditbookmarks: Manage Bookmarks functionality'
            'kde-cli-tools: manage file associations, root mode'
            'kompare: file contents comparison'
            'kdiff3: file contents comparison'
            'krename: advanced rename tool'
            'konsolepart: terminal'
            'ktexteditor: file editing support')
source=(https://download.kde.org/stable/$pkgname/$pkgver/$pkgname-$pkgver.tar.xz{,.sig})
sha256sums=('c9b79bfade6cc69fe0e341ecef932fcac8afd9fe94e8cbcfbd729feb54394e04'
            'SKIP')
validpgpkeys=(3381929FB59A539333042D57AD596556DFF2501F  # Nikita Melnichenko <nikita+kde@melnichenko.name>
              03D30A1182C9000A769666A250DA16FAE2CC4E5B) # A. Bikadorov <alex.bikadorov@kdemail.net>

build() {
  cmake -B build -S $pkgname-$pkgver \
    -DKDESU_PATH="/usr/lib/kf6/kdesu"
  cmake --build build
}

package() {
  DESTDIR="$pkgdir" cmake --install build
}
