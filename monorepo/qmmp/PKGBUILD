# Maintainer: Jaroslav Lichtblau <svetlemodry@archlinux.org>

pkgname=qmmp
pkgver=2.4.2
pkgrel=1
pkgdesc="Qt based audio-player"
arch=('x86_64')
url="http://qmmp.ylsoftware.com/"
license=('GPL')
depends=('alsa-lib' 'curl' 'desktop-file-utils' 'hicolor-icon-theme' 'libmad' 'libvorbis'
         'libogg' 'taglib' 'xdg-utils' 'qt6-base' 'qt6-multimedia')
makedepends=('cmake' 'flac' 'jack' 'libmpcdec' 'pipewire' 'pulseaudio' 'ffmpeg' 'libcdio-paranoia'
             'libcddb' 'libmms' 'libsamplerate' 'libmodplug' 'libsndfile' 'wavpack' 'mpg123' 'mesa'
             'projectm' 'faad2' 'libgme' 'libsidplayfp' 'opusfile>=0.12' 'wildmidi' 'qt6-tools')
optdepends=('mpg123: to play MP3 files'
            'flac: native FLAC support'
            'jack: JACK sound output'
            'libmpcdec: Musepack support'
            'pulseaudio: PulseAudio output'
            'pipewire: pipewire support'
            'libcdio-paranoia: Compact Disc input and control support'
            'libcddb: CDDB server support'
            'libmms: MMS stream protocol support'
            'libmodplug: MOD playing library'
            'libsndfile: sampled sound support'
            'projectm: visual efects'
            'faad2: ADTS AAC support'
            'libgme: support for chiptunes from various platforms'
            'libsidplayfp: C64 music support'
            'opusfile: Opus support'
            'wildmidi: MIDI support'
            'ffmpeg' 'libsamplerate' 'wavpack')
changelog=$pkgname.changelog
source=(http://qmmp.ylsoftware.com/files/$pkgname/${pkgver%.*}/$pkgname-$pkgver.tar.bz2)
sha512sums=('1d20fc98960c3b5482bf8688a42c7788ed0e5a8b3b587adf7b7f7134403df7504eb4ac22412475ee4e3153e0a1fb09be24d4959317902794a44b687addbe37c2')

build() {
  cd "${srcdir}"/$pkgname-$pkgver

  cmake . -DCMAKE_INSTALL_PREFIX=/usr
  make
}

package() {
  cd "${srcdir}"/$pkgname-$pkgver

  make DESTDIR="${pkgdir}" install
}
